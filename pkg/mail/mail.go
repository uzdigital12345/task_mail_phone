package mail

import (
	"log"
	"net/smtp"
)

const (
	from = "email@gmail.com"
	pass = "password"
	to   = "email@gmail.com"
)

type Mail struct {
	body string
}

func New(body string) *Mail {
	return &Mail{
		body: body,
	}
}

func (m Mail) SendToMail(body string) error {

	message := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Assolomu alekum!!!\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(message))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}
	return nil
}
