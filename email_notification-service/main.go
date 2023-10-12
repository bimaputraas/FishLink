package main

import (
	"final_project-ftgo-h8/config"
	"final_project-ftgo-h8/email_notification-service/consumer"
	"final_project-ftgo-h8/helper"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main(){
	// load env
	helper.LoadEnv()

	// init channel
	channel := config.NewChannel()
	
	// init queue
	registerQueue := config.AddQueue(channel,"fishlink-email_notification")
	
	// init consumer
	registerQonsumer := consumer.NewRegisterNotification(channel)
	
	// start register consume
	go registerQonsumer.ConsumeQueuedMessage(registerQueue.Name)
	
	// continue
	var forever chan struct{}
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}