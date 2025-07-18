package mail

import (
	"fmt"
	mail2 "github.com/pascalallen/carline/internal/carline/domain/mail"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridMailService struct {
	client *sendgrid.Client
}

func NewSendGridMailService(client *sendgrid.Client) mail2.Service {
	return &SendGridMailService{client: client}
}

func (s *SendGridMailService) Send(from mail2.Sender, to mail2.Recipient, message mail2.Message) error {
	f := mail.NewEmail(from.Name, from.EmailAddress)
	t := mail.NewEmail(to.Name, to.EmailAddress)
	msg := mail.NewSingleEmail(f, message.Subject, t, message.PlainTextBody, message.HtmlBody)

	res, err := s.client.Send(msg)
	if err != nil {
		return fmt.Errorf("error sending mail message via SendGrid. error: %v, response: %v", err, res)
	}

	return nil
}
