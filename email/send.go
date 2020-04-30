package email

import (
	"log"
	"net/smtp"
	"os"
	"sync"
)

// SendEmail Function sends the email.
func SendEmail(body, to string, wg *sync.WaitGroup) {
	defer wg.Done()
	from := os.Getenv("FROM")
	pass := os.Getenv("PASS")
	// to := os.Getenv("TO")
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Google Testing\n" +
		body
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	log.Printf("sent, visit to gmail [%s]", to)
}
