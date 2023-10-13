package consumer

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer interface {
	ConsumeQueuedMessage(queueName string)
}
type registerNotification struct {
	channel *amqp.Channel
}

func NewRegisterNotification(c *amqp.Channel) Consumer{
	return &registerNotification{
		channel: c,
	}
}
