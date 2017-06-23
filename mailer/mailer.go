package mailer

import (
	"strconv"

	"github.com/gobuffalo/envy"
	gomail "gopkg.in/gomail.v2"
)

//Mailer is an interface for different types of mailers each one will implement needed logics.
type Mailer interface {
	Send(m Message) error
}

type smtpMailer struct {
	dialer *gomail.Dialer
}

//Send a message using SMTP configuration or returns an error if something goes wrong.
func (sm smtpMailer) Send(m Message) error {
	return nil
}

//NewSMTPMailer Creates an SMTP mailer by reading configuration from the env or using defaults.
func NewSMTPMailer() smtpMailer {
	port, _ := strconv.Atoi(envy.Get("SMTP_PORT", "1025"))
	dialer := gomail.NewDialer(envy.Get("SMTP_HOST", "localhost"), port, envy.Get("SMTP_USER", ""), envy.Get("SMTP_PASSWORD", ""))

	return smtpMailer{
		dialer: dialer,
	}
}
