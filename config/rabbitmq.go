package config

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)
func NewChannel() *amqp.Channel{
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err.Error())
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err.Error())
	}

	return ch
}

func AddQueue(ch *amqp.Channel,queueName string) amqp.Queue {
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	return q
}