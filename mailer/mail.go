package mailer

import (
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/plush"
	"github.com/pkg/errors"
)

var templateBox = packr.NewBox("../templates")

//Message represents an Email message
type Message struct {
	From    string
	To      string
	Subject string
	Body    string
}

//BuildBody builds the message body based on passed template and mail data.
func (m Message) BuildBody(templatePath string, data map[string]interface{}) error {
	tmpl, err := templateBox.MustString(templatePath)
	if err != nil {
		return errors.WithStack(err)
	}

	s, err := plush.Render(tmpl, plush.NewContextWith(data))
	if err != nil {
		return errors.WithStack(err)
	}

	m.Body = s
	return nil
}
