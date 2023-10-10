package consumer

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type emailNotification struct {
	channel *amqp.Channel
}

func NewConsumer(c *amqp.Channel) Consumer{
	return &emailNotification{channel: c}
}
