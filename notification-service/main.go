package main

import (
	"fishlink-notification-service/config"
	"fishlink-notification-service/consumer"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main(){
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