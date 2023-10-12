package config

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)
func NewChannel() *amqp.Channel{
	rabbitmqHost := os.Getenv("RABBITMQ_HOST")
	conn, err := amqp.Dial("amqp://guest:guest@"+rabbitmqHost+":5672/")
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