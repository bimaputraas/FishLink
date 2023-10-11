package consumer

import (
	"encoding/json"
	"final_project-ftgo-h8/api/dto"
	"final_project-ftgo-h8/helper"
	"fmt"
	"log"
)

func (c *emailNotification) ConsumeQueuedMessage(){
	msgs, err := c.channel.Consume(
		c.queue.Name, // queue
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
	  
	var forever chan struct{}
	
	go func() {
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
	}()
	  
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
