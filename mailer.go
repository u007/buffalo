package buffalo

import (
	"strconv"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/plush"
	"github.com/pkg/errors"
	gomail "gopkg.in/gomail.v2"
)

//EmailData is the data needed to send the email.
type EmailData struct {
	From    string
	To      string
	Subject string
	//Other fields that we should add here.
	TemplateData map[string]interface{}
}

//BuildBody builds the message body based on pased data and the template.
func (m EmailData) BuildBody(templatesBox packr.Box, templatePath string, data EmailData) (string, error) {
	tmpl, err := templatesBox.MustString(templatePath)
	if err != nil {
		return "", err
	}

	s, err := plush.Render(tmpl, plush.NewContextWith(data.TemplateData))
	if err != nil {
		return "", err
	}

	return s, nil
}

//Mailer is an interface for different types of mailers each one will implement needed logics.
type Mailer interface {
	Send(m EmailData, templatePath string, contentType string) error
}

//SMTPMailer is the first implementation of the mailer interface.
type SMTPMailer struct {
	Dialer      *gomail.Dialer
	TemplateBox packr.Box
}

//Send a message using SMTP configuration or returns an error if something goes wrong.
func (sm SMTPMailer) Send(data EmailData, templatePath string, contentType string) error {
	message := gomail.NewMessage()

	message.SetHeader("From", data.From)
	message.SetHeader("To", data.To)
	message.SetHeader("Subject", data.Subject)

	body, err := data.BuildBody(sm.TemplateBox, templatePath, data)
	if err != nil {
		return errors.WithStack(err)
	}
	message.SetBody(contentType, body)

	err = sm.Dialer.DialAndSend(message)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

//NewSMTPMailer Creates an SMTP mailer by reading configuration from the env or using defaults.
func NewSMTPMailer(box packr.Box) SMTPMailer {
	port, _ := strconv.Atoi(envy.Get("SMTP_PORT", "1025"))
	dialer := gomail.NewDialer(envy.Get("SMTP_HOST", "localhost"), port, envy.Get("SMTP_USER", ""), envy.Get("SMTP_PASSWORD", ""))

	return SMTPMailer{
		Dialer:      dialer,
		TemplateBox: box,
	}
}
