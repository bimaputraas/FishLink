package publisher

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (pub *emailNotification) PublishMessage(ctx context.Context,queueName string, message []byte) error{
	err := pub.channel.PublishWithContext(ctx,
	"",     // exchange
	queueName, // routing key
	false,  // mandatory
	false,  // immediate
	amqp.Publishing {
		ContentType: "text/plain",
		Body:        message,
	})
	if err != nil {
		return err
	}
	
	return nil
}

