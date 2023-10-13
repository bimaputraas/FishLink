package consumer

import (
	"encoding/json"
	"fishlink-notification-service/dto"
	"fishlink-notification-service/helper"
	"fmt"
	"log"
)

func (c *registerNotification) ConsumeQueuedMessage(queueName string){
	msgs, err := c.channel.Consume(
		queueName, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal(err)
	}
	
	
	for d := range msgs {
	  	log.Printf("Received a message: %s", d.Body)
		  
	  	var userVerif dto.UserEmailVerification
	  	err := json.Unmarshal(d.Body, &userVerif)
	  	if err != nil {
				log.Fatal(err)
	  	}

	  	subject := "Fishlink account verification"
	  	message := fmt.Sprintf("Your verification link : http://localhost:8080/user-verification-register/%d/%s", userVerif.UserId,userVerif.VerificationCode)

		helper.SendMail(userVerif.Email, subject, message)
	}
	
}
