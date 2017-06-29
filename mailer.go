package buffalo

import (
	"strconv"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr"
	gomail "gopkg.in/gomail.v2"
)

//EmailData is the args passed to render.
type EmailData map[string]interface{}

//Mailer is an interface for different types of mailers each one will implement needed logics.
type Mailer interface {
	Send(m EmailData) error
}

//SMTPMailer is the first implementation of the mailer interface.
type SMTPMailer struct {
	Dialer      *gomail.Dialer
	TemplateBox packr.Box
}

//Send a message using SMTP configuration or returns an error if something goes wrong.
func (sm SMTPMailer) Send(data EmailData) error {
	return nil
}

//NewSMTPMailer Creates an SMTP mailer by reading configuration from the env or using defaults.
func NewSMTPMailer(box packr.Box) *SMTPMailer {
	port, _ := strconv.Atoi(envy.Get("SMTP_PORT", "1025"))
	dialer := gomail.NewDialer(envy.Get("SMTP_HOST", "localhost"), port, envy.Get("SMTP_USER", ""), envy.Get("SMTP_PASSWORD", ""))

	return &SMTPMailer{
		Dialer:      dialer,
		TemplateBox: box,
	}
}
