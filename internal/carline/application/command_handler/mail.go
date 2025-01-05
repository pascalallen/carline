package command_handler

import (
	"bytes"
	"fmt"
	"github.com/pascalallen/carline/internal/carline/application/command"
	"github.com/pascalallen/carline/internal/carline/application/event"
	"github.com/pascalallen/carline/internal/carline/domain/mail"
	"github.com/pascalallen/carline/internal/carline/domain/security_token"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"html/template"
	"os"
	"time"
)

type SendWelcomeEmailHandler struct {
	SecurityTokenService security_token.Service
	EventDispatcher      messaging.EventDispatcher
	MailService          mail.Service
}

func (h SendWelcomeEmailHandler) Handle(cmd messaging.Command) error {
	c, ok := cmd.(*command.SendWelcomeEmail)
	if !ok {
		return fmt.Errorf("invalid command type passed to SendWelcomeEmailHandler: %v", cmd)
	}

	activationToken, err := h.SecurityTokenService.FetchById(c.SecurityTokenId)
	if err != nil {
		return fmt.Errorf("error retrieving activation token: %s", err)
	}

	data := struct {
		Subject   string
		FirstName string
		BaseUrl   string
		Token     string
		Year      int
	}{
		Subject:   "Welcome to Carline!",
		FirstName: c.FirstName,
		BaseUrl:   os.Getenv("APP_BASE_URL"),
		Token:     string(activationToken.Crypto),
		Year:      time.Now().Year(),
	}

	tmpl, err := template.ParseFiles("web/template/auth/activation.tmpl")
	if err != nil {
		return fmt.Errorf("error parsing activation template: %s", err)
	}

	var tplBuffer bytes.Buffer
	if err := tmpl.Execute(&tplBuffer, data); err != nil {
		return fmt.Errorf("error executing activation template: %s", err)
	}

	htmlContent := tplBuffer.String()

	from := mail.Sender{
		Name:         "Carline Team",
		EmailAddress: "postmaster@mg.pascalallen.com",
	}
	to := mail.Recipient{
		Name:         c.FirstName + " " + c.LastName,
		EmailAddress: c.EmailAddress,
	}
	msg := mail.Message{
		Subject:       "Welcome to Carline!",
		PlainTextBody: "", // TODO
		HtmlBody:      htmlContent,
	}

	err = h.MailService.Send(from, to, msg)
	if err != nil {
		return fmt.Errorf("error sending welcome email: %s", err)
	}

	evt := event.WelcomeEmailSent{
		Id:              c.Id,
		FirstName:       c.FirstName,
		LastName:        c.LastName,
		EmailAddress:    c.EmailAddress,
		SecurityTokenId: c.SecurityTokenId,
	}
	h.EventDispatcher.Dispatch(evt)

	return nil
}
