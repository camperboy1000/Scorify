package minion

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/checks"
	"github.com/scorify/scorify/pkg/config"
	"github.com/scorify/scorify/pkg/ent/minion"
	"github.com/scorify/scorify/pkg/ent/status"
	"github.com/scorify/scorify/pkg/rabbitmq/rabbitmq"
	"github.com/scorify/scorify/pkg/structs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "minion",
	Short:   "Start scoring minion worker",
	Long:    "Start scoring minion worker",
	Aliases: []string{"m", "worker", "w"},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.InitMinion()
	},

	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	for {
		logrus.Info("Starting minion worker")
		mainLoop()
		logrus.Info("minion worker stopped")
		time.Sleep(time.Second * 5)
	}
}

func mainLoop() {
	rabbitmqClient, err := rabbitmq.Client(
		config.RabbitMQ.Minion.User,
		config.RabbitMQ.Minion.Password,
	)
	if err != nil {
		logrus.WithError(err).Error("failed to create rabbitmq client")
		return
	}
	defer rabbitmqClient.Close()

	backOff := time.Second
	heartbeatSuccess := make(chan struct{})
	ctx := context.Background()

	// Reset backoff on successful heartbeat
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-heartbeatSuccess:
				backOff = time.Second
			}
		}
	}()

	// Run minion loop first with no backoff first
	minionLoop(context.Background(), rabbitmqClient, heartbeatSuccess)
	for {
		time.Sleep(backOff)
		minionLoop(context.Background(), rabbitmqClient, heartbeatSuccess)
		backOff = min(backOff*2, time.Minute)
	}
}

func minionLoop(ctx context.Context, rabbitmqClient *rabbitmq.RabbitMQConnections, heartbeatSuccess chan struct{}) {
	minionCtx, minionCancel := context.WithCancel(ctx)
	defer minionCancel()

	workerEnrollClient, err := rabbitmqClient.WorkerEnrollClient()
	if err != nil {
		logrus.WithError(err).Fatal("failed to create worker enroll client")
	}

	err = workerEnrollClient.EnrollMinion(minionCtx, minion.RoleService)
	if err != nil {
		logrus.WithError(err).Fatal("failed to enroll minion")
	}
	workerEnrollClient.Close()

	heartbeatClient, err := rabbitmqClient.HeartbeatClient()
	if err != nil {
		logrus.WithError(err).Fatal("failed to create heartbeat client")
	}
	defer heartbeatClient.Close()

	workerStatusListener, err := rabbitmqClient.WorkerStatusListener(ctx)
	if err != nil {
		logrus.WithError(err).Fatal("failed to create worker status listener")
	}
	defer workerStatusListener.Close()

	workerStatusChannel := workerStatusListener.Consume(minionCtx)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	err = heartbeatClient.SendHeartbeat(minionCtx)
	if err != nil {
		logrus.WithError(err).Error("failed to send heartbeat")
		return
	}

	heartbeatSuccess <- struct{}{}

	scoringCtx, scoringCtxCancel := context.WithCancel(minionCtx)
	defer scoringCtxCancel()

	scoring := true

	go score(scoringCtx, rabbitmqClient)

	for {
		select {
		case <-minionCtx.Done():
			return
		case <-ticker.C:
			err := heartbeatClient.SendHeartbeat(minionCtx)
			if err != nil {
				logrus.WithError(err).Error("failed to send heartbeat")
				return
			}

			select {
			case <-time.After(5 * time.Second):
				logrus.Error("failed to publish to heartbeatSuccess channel")
				return
			case <-minionCtx.Done():
				return
			case heartbeatSuccess <- struct{}{}:
			}

		case workerStatus := <-workerStatusChannel:
			disabled := workerStatus.Disabled(config.Minion.ID)

			if disabled && scoring {
				scoringCtxCancel()
			} else if !disabled && !scoring {
				go score(scoringCtx, rabbitmqClient)
			}

		case <-scoringCtx.Done():
			scoring = true
		}
	}
}

func score(ctx context.Context, rabbitmqClient *rabbitmq.RabbitMQConnections) {
	taskRequestListener, err := rabbitmqClient.TaskRequestListener(ctx)
	if err != nil {
		logrus.WithError(err).Fatal("failed to create task request listener")
	}
	defer taskRequestListener.Close()

	taskResponseClient, err := rabbitmqClient.TaskResponseClient()
	if err != nil {
		logrus.WithError(err).Fatal("failed to create task response client")
	}
	defer taskResponseClient.Close()

	for {
		// recieved score task
		task, err := taskRequestListener.Consume(ctx)
		if err != nil {
			logrus.WithError(err).Error("failed to consume task request")
			return
		}

		checkDeadline := time.Now().Add(time.Duration(float64(config.Interval) * 0.8))
		submissionDeadline := time.Now().Add(time.Duration(float64(config.Interval) * 0.85))

		// check if source exists
		check, ok := checks.Checks[task.SourceName]
		if !ok {
			logrus.WithField("source", task.SourceName).
				Error("source not found")

			err = taskResponseClient.SubmitTaskResponse(
				ctx,
				&structs.TaskResponse{
					StatusID: task.StatusID,
					MinionID: config.Minion.ID,
					Status:   status.StatusDown,
					Error:    "source not found",
				},
			)
			if err != nil {
				logrus.WithError(err).Error("encountered error while submitting score task")
				ctx.Done()
				return
			}
			continue
		}

		// run check
		go func(checkDeadline time.Time, submissionDeadline time.Time, uuid uuid.UUID, check checks.Check, task_config string) {
			checkCtx, checkCancel := context.WithDeadline(context.Background(), checkDeadline)
			submissionCtx, submissionCancel := context.WithDeadline(context.Background(), submissionDeadline)
			defer checkCancel()
			defer submissionCancel()

			// run check and close check context
			err := check.Func(checkCtx, task_config)
			checkCtx.Done()

			// submit score task and close submission context
			if err != nil {
				err = taskResponseClient.SubmitTaskResponse(
					submissionCtx,
					&structs.TaskResponse{
						StatusID: task.StatusID,
						MinionID: config.Minion.ID,
						Status:   status.StatusDown,
						Error:    err.Error(),
					},
				)
			} else {
				err = taskResponseClient.SubmitTaskResponse(
					submissionCtx,
					&structs.TaskResponse{
						StatusID: task.StatusID,
						MinionID: config.Minion.ID,
						Status:   status.StatusUp,
					},
				)
			}
			submissionCtx.Done()

			// log error if submission failed
			if err != nil {
				logrus.WithError(err).Error("encountered error while submitting score task")
			}
		}(checkDeadline, submissionDeadline, task.StatusID, check, task.Config)
	}
}
