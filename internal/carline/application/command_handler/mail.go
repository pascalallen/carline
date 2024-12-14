package command_handler

import (
	"fmt"
	"github.com/pascalallen/carline/internal/carline/application/command"
	"github.com/pascalallen/carline/internal/carline/application/event"
	"github.com/pascalallen/carline/internal/carline/domain/mail"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

type SendWelcomeEmailHandler struct {
	EventDispatcher messaging.EventDispatcher
	MailService     mail.Service
}

func (h SendWelcomeEmailHandler) Handle(cmd messaging.Command) error {
	c, ok := cmd.(*command.SendWelcomeEmail)
	if !ok {
		return fmt.Errorf("invalid command type passed to SendWelcomeEmailHandler: %v", cmd)
	}

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
		PlainTextBody: "Please follow this link to activate your account",
	}
	err := h.MailService.Send(from, to, msg)
	if err != nil {
		return err
	}

	evt := event.WelcomeEmailSent{
		Id:           c.Id,
		FirstName:    c.FirstName,
		LastName:     c.LastName,
		EmailAddress: c.EmailAddress,
		Token:        c.Token,
	}
	h.EventDispatcher.Dispatch(evt)

	return nil
}
