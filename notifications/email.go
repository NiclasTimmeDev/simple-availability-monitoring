package notifications

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmailNotification(url string, body string) {

	if os.Getenv("EMAILS_ENABLED") == "" {
		return
	}
	
	// For authentication.
	host := os.Getenv("MAIL_HOST")
	port := os.Getenv("MAIL_PORT")
	address := host + ":" + port
	user := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")

	// Mail data.
	from := os.Getenv("MAIL_FROM_ADDRESS")
	to := os.Getenv("MAIL_TO_ADDRESS")
	subject := fmt.Sprintf("Subject: Monitoring - Error while checking %s\n\n", url)
	msg := []byte(subject + body)

	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(address, auth, from, []string{to}, []byte(msg))
	if err != nil {
		panic(err)
	}

}