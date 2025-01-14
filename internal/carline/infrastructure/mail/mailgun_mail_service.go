package mail

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/pascalallen/carline/internal/carline/domain/mail"
	"time"
)

type MailgunMailService struct {
	client *mailgun.MailgunImpl
}

func NewMailgunMailService(client *mailgun.MailgunImpl) mail.Service {
	return &MailgunMailService{client: client}
}

func (m *MailgunMailService) Send(from mail.Sender, to mail.Recipient, message mail.Message) error {
	msg := mailgun.NewMessage(from.EmailAddress, message.Subject, "", to.EmailAddress)
	msg.SetHTML(message.HtmlBody)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, err := m.client.Send(ctx, msg)
	if err != nil {
		return fmt.Errorf("error sending mail message via Mailgun: %v", err)
	}

	return nil
}
