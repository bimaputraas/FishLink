package publisher

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher interface {
	PublishMessage(ctx context.Context,queueName string, msg []byte) error
}

type emailNotification struct {
	channel *amqp.Channel
}

func NewPublisher(c *amqp.Channel) Publisher{
	return &emailNotification{channel: c}
}
