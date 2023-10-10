package publisher

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type emailNotification struct {
	channel *amqp.Channel
}

func NewPublisher(c *amqp.Channel) Publisher{
	return &emailNotification{channel: c}
}
