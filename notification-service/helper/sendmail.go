package helper

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "FishLink <ddummymail65@gmail.com>"
const CONFIG_AUTH_EMAIL = "ddummymail65@gmail.com"

func SendMail(mail, subject, message string) error {
	// auth password
	var CONFIG_AUTH_PASSWORD = os.Getenv("AUTHMAILPASSWORD")
	// body
	to := []string{mail}
	cc := []string{}
	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}
	log.Printf("success send mail to %s", mail)

	return nil
}