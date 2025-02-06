package mail

import (
	"fmt"
	"github.com/pascalallen/carline/internal/carline/domain/mail"
	"net/smtp"
)

type MailtrapMailService struct {
	Address string
	Auth    smtp.Auth
}

func NewMailtrapMailService(host string, port string, username string, password string) mail.Service {
	addr := fmt.Sprintf("%s:%s", host, port)
	auth := smtp.PlainAuth("", username, password, host)
	return &MailtrapMailService{
		Address: addr,
		Auth:    auth,
	}
}

func (m *MailtrapMailService) Send(from mail.Sender, to mail.Recipient, message mail.Message) error {
	msg := []byte(fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/html; charset=\"UTF-8\"\r\n"+
			"\r\n"+
			"%s",
		from.EmailAddress,
		to.EmailAddress,
		message.Subject,
		message.HtmlBody,
	))
	return smtp.SendMail(m.Address, m.Auth, from.EmailAddress, []string{to.EmailAddress}, msg)
}
