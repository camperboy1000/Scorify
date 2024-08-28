package rabbitmq

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/scorify/scorify/pkg/rabbitmq/types"
)

const (
	TaskResponseQueue = "task_response_queue"
	TaskResponseVhost = "task_response_vhost"
)

var (
	// Permissions for minions in task_response vhost
	TaskResponseConfigurePermissions   = regex(TaskResponseQueue)
	TaskResponseMinionWritePermissions = regex_amq_default(TaskResponseQueue)
	TaskResponseMinionReadPermissions  = regex("")
)

func taskResponseQueue(conn *amqp.Connection) (*amqp.Channel, amqp.Queue, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	q, err := ch.QueueDeclare(
		TaskResponseQueue,
		false,
		false,
		false,
		false,
		nil,
	)

	return ch, q, err
}

type taskResponseListener struct {
	ch   *amqp.Channel
	msgs <-chan amqp.Delivery
}

func TaskResponseListener(ctx context.Context, conn *amqp.Connection) (*taskResponseListener, error) {
	ch, q, err := taskResponseQueue(conn)
	if err != nil {
		return nil, err
	}

	msgs, err := ch.ConsumeWithContext(
		ctx,
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &taskResponseListener{
		ch:   ch,
		msgs: msgs,
	}, nil
}

func (l *taskResponseListener) Close() error {
	return l.ch.Close()
}

func (l *taskResponseListener) Consume(ctx context.Context) (*types.TaskResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case msg := <-l.msgs:
		var taskResponse types.TaskResponse
		err := json.Unmarshal(msg.Body, &taskResponse)
		if err != nil {
			return nil, err
		}

		return &taskResponse, nil
	}
}

type taskResponseClient struct {
	ch *amqp.Channel
	q  amqp.Queue
}

func TaskResponseClient(conn *amqp.Connection, ctx context.Context) (*taskResponseClient, error) {
	ch, q, err := taskResponseQueue(conn)
	if err != nil {
		return nil, err
	}

	return &taskResponseClient{
		ch: ch,
		q:  q,
	}, nil
}

func (c *taskResponseClient) Close() error {
	return c.ch.Close()
}

func (c *taskResponseClient) SubmitTaskResponse(ctx context.Context, taskResponse *types.TaskResponse) error {
	out, err := json.Marshal(taskResponse)
	if err != nil {
		return err
	}

	return c.ch.PublishWithContext(
		ctx,
		"",
		c.q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        out,
		},
	)
}
