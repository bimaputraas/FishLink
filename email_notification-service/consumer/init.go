package consumer

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer interface {
	ConsumeQueuedMessage()
}
type emailNotification struct {
	channel *amqp.Channel
	queue amqp.Queue
}

func NewConsumer(c *amqp.Channel, q amqp.Queue) Consumer{
	return &emailNotification{
		channel: c,
		queue: q,
	}
}
