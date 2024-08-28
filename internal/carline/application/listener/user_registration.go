package listener

import (
	"fmt"
	"github.com/pascalallen/carline/internal/carline/application/command"
	"github.com/pascalallen/carline/internal/carline/application/event"
	"github.com/pascalallen/carline/internal/carline/domain/crypto"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

type UserRegistration struct {
	CommandBus messaging.CommandBus
}

func (l UserRegistration) Handle(evt messaging.Event) error {
	e, ok := evt.(*event.UserRegistered)
	if !ok {
		return fmt.Errorf("invalid event type passed to UserRegistration listener: %v", evt)
	}

	token := crypto.Generate()
	cmd := command.SendWelcomeEmail{
		Id:           e.Id,
		FirstName:    e.FirstName,
		LastName:     e.LastName,
		EmailAddress: e.EmailAddress,
		Token:        token,
	}
	l.CommandBus.Execute(cmd)

	return nil
}
