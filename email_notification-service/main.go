package main

import (
	"final_project-ftgo-h8/config"
	"final_project-ftgo-h8/email_notification-service/consumer"

	_ "github.com/joho/godotenv/autoload"
)

func main(){
	// init channel
	channel := config.NewChannel()

	// add queue
	queue := config.AddQueue(channel,"fishlink-email_notification")

	// init consumer
	consumer := consumer.NewConsumer(channel)

	// start app
	consumer.ConsumeQueuedMessage(queue.Name)
}