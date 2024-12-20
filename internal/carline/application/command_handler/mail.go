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
)

type SendWelcomeEmailHandler struct {
	SecurityTokenRepository security_token.Repository
	EventDispatcher         messaging.EventDispatcher
	MailService             mail.Service
}

func (h SendWelcomeEmailHandler) Handle(cmd messaging.Command) error {
	c, ok := cmd.(*command.SendWelcomeEmail)
	if !ok {
		return fmt.Errorf("invalid command type passed to SendWelcomeEmailHandler: %v", cmd)
	}

	activationToken, err := h.SecurityTokenRepository.GetById(c.SecurityTokenId)
	if err != nil {
		return fmt.Errorf("error retrieving activation token: %s", err)
	}

	data := struct {
		FirstName string
		BaseUrl   string
		Token     string
	}{
		FirstName: c.FirstName,
		BaseUrl:   os.Getenv("APP_BASE_URL"),
		Token:     string(activationToken.Crypto),
	}

	tmpl, err := template.ParseFiles("activation.tmpl")
	if err != nil {
		return fmt.Errorf("error parsing activation template: %s", err)
	}

	var tplBuffer bytes.Buffer
	if err := tmpl.Execute(&tplBuffer, data); err != nil {
		return fmt.Errorf("error executing activation template: %s", err)
	}

	htmlContent := tplBuffer.String()

	from := mail.Sender{
		Name:         "Pascal Allen",
		EmailAddress: "pascal.allen88@gmail.com",
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
		return err
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
