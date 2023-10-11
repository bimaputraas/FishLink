package main

import (
	"final_project-ftgo-h8/config"
	"final_project-ftgo-h8/email_notification-service/consumer"
	"final_project-ftgo-h8/helper"

	_ "github.com/joho/godotenv/autoload"
)

func main(){
	// load env
	helper.LoadEnv()

	// init channel
	channel := config.NewChannel()
	
	// init queue
	queue := config.AddQueue(channel,"fishlink-email_notification")
	
	// init consumer
	consumer := consumer.NewConsumer(channel,queue)
	
	// start app
	consumer.ConsumeQueuedMessage()
}