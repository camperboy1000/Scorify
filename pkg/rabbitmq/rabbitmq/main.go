package rabbitmq

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/scorify/scorify/pkg/config"
	"github.com/scorify/scorify/pkg/rabbitmq/management/types"
	"github.com/sirupsen/logrus"
)

type rabbitMQConnections struct {
	Heartbeat    *amqp.Connection
	TaskRequest  *amqp.Connection
	TaskResponse *amqp.Connection
	WorkerStatus *amqp.Connection
}

func (r *rabbitMQConnections) Close() error {
	err := r.Heartbeat.Close()
	if err != nil {
		return err
	}

	err = r.TaskRequest.Close()
	if err != nil {
		return err
	}

	err = r.TaskResponse.Close()
	if err != nil {
		return err
	}

	return r.WorkerStatus.Close()
}

func openConnection(vhost string, username string, password string) (*amqp.Connection, error) {
	connStr := fmt.Sprintf(
		"amqp://%s:%s@%s:%d/%s",
		username,
		password,
		config.RabbitMQ.Host,
		config.RabbitMQ.Port,
		vhost,
	)

	conn, err := amqp.Dial(connStr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func openClient(username string, password string) (*rabbitMQConnections, error) {
	heartbeatConn, err := openConnection(HeartbeatVhost, username, password)
	if err != nil {
		return nil, err
	}

	taskRequestsConn, err := openConnection(TaskRequestVhost, username, password)
	if err != nil {
		return nil, err
	}

	taskResponsesConn, err := openConnection(TaskResponseVhost, username, password)
	if err != nil {
		return nil, err
	}

	workerStatusConn, err := openConnection(WorkerStatusVhost, username, password)
	if err != nil {
		return nil, err
	}

	return &rabbitMQConnections{
		Heartbeat:    heartbeatConn,
		TaskRequest:  taskRequestsConn,
		TaskResponse: taskResponsesConn,
		WorkerStatus: workerStatusConn,
	}, nil
}

func Serve(ctx context.Context) error {
	conn, err := openClient(config.RabbitMQ.Server.User, config.RabbitMQ.Server.Password)
	if err != nil {
		return err
	}
	defer conn.Close()

	logrus.Info("Connected to RabbitMQ server")

	go func() {
		for {
			err := ListenHeartbeat(
				conn.Heartbeat,
				func(heartbeat *types.Heartbeat) {
					fmt.Println("Received heartbeat: ", heartbeat)
				},
				ctx,
			)
			if err != nil {
				logrus.WithError(err).Error("encountered error while listening to heartbeats")
			}
		}
	}()

	go func() {
		taskRequestClient, err := TaskRequestClient(conn.TaskRequest, ctx)
		if err != nil {
			logrus.WithError(err).Fatal("failed to create task request client")
		}

		ticker := time.NewTicker(time.Second)

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				err := taskRequestClient.Publish(ctx, &types.TaskRequest{})
				if err != nil {
					logrus.WithError(err).Fatal("failed to send task request")
				}
			}
		}
	}()

	go func() {
		for {
			err := ListenTaskResponse(
				conn.TaskResponse,
				func(taskResponse *types.TaskResponse) {
					fmt.Println("Received task response: ", taskResponse)
				},
				ctx,
			)
			if err != nil {
				logrus.WithError(err).Fatal("failed to listen to task response")
			}
		}
	}()

	go func() {
		workerStatusClient, err := WorkerStatusClient(conn.WorkerStatus)
		if err != nil {
			logrus.WithError(err).Fatal("failed to create worker status client")
		}

		ticker := time.NewTicker(time.Second)

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				fmt.Println("sending worker status")
				err := workerStatusClient.Publish(ctx, &types.WorkerStatus{})
				if err != nil {
					logrus.WithError(err).Fatal("failed to send worker status")
				}
			}
		}
	}()

	<-ctx.Done()

	return nil
}

func Client(ctx context.Context) error {
	conn, err := openClient(config.RabbitMQ.Minion.User, config.RabbitMQ.Minion.Password)
	if err != nil {
		return err
	}

	logrus.Info("RabbitMQ client started")

	go func() {
		heartbeatClient, err := HeartbeatClient(conn.Heartbeat, ctx)
		if err != nil {
			logrus.WithError(err).Error("failed to create heartbeat client")
		}

		ticker := time.NewTicker(time.Second)

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				err := heartbeatClient.SendHeartbeat(ctx)
				if err != nil {
					logrus.WithError(err).Error("failed to send heartbeat")
				}
			}
		}
	}()

	go func() {
		for {
			err := ListenTaskRequest(
				conn.TaskRequest,
				func(taskRequest *types.TaskRequest) {
					fmt.Println("Received task request: ", taskRequest)
				},
				ctx,
			)
			if err != nil {
				logrus.WithError(err).Error("failed to listen task request")
			}
		}
	}()

	go func() {
		taskResponseClient, err := TaskResponseClient(conn.TaskResponse, ctx)
		if err != nil {
			logrus.WithError(err).Error("failed to create task response client")
		}

		ticker := time.NewTicker(time.Second)

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				err := taskResponseClient.SubmitTaskResponse(ctx, &types.TaskResponse{})
				if err != nil {
					logrus.WithError(err).Error("failed to send task response")
				}
			}
		}
	}()

	go func() {
		err := ListenWorkerStatus(
			conn.WorkerStatus,
			func(workerStatus *types.WorkerStatus) {
				fmt.Println("Received worker status: ", workerStatus)
			},
			ctx,
		)
		if err != nil {
			logrus.WithError(err).Error("failed to listen worker status")
		}
	}()

	<-ctx.Done()

	return nil
}
