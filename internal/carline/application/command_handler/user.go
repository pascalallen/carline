package command_handler

import (
	"fmt"
	"github.com/pascalallen/carline/internal/carline/application/command"
	"github.com/pascalallen/carline/internal/carline/domain/event"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/storage"
)

type ProjectionState struct {
	EmailAddresses map[string]string `json:"email_addresses"`
}

type UpdateUserEmailAddressHandler struct {
	EventStore storage.EventStore
}

func (h UpdateUserEmailAddressHandler) Handle(cmd messaging.Command) error {
	c, ok := cmd.(*command.UpdateUserEmailAddress)
	if !ok {
		return fmt.Errorf("invalid command type passed to UpdateUserEmailAddressHandler: %v", cmd)
	}

	var result ProjectionState
	if err := h.EventStore.UnmarshalProjectionResult("user-email-addresses", &result); err != nil {
		return fmt.Errorf("error getting projection result: %v", err)
	}

	for emailAddress := range result.EmailAddresses {
		if emailAddress == c.EmailAddress {
			return fmt.Errorf("could not update user. email address %s is already taken", c.EmailAddress)
		}
	}

	emailUpdateEvent := event.UserEmailAddressUpdated{
		Id:           c.Id,
		EmailAddress: c.EmailAddress,
	}
	streamId := fmt.Sprintf("user-%s", c.Id)
	err := h.EventStore.AppendToStream(streamId, emailUpdateEvent)
	if err != nil {
		return fmt.Errorf("could not store UserEmailAddressUpdated event: %w", err)
	}

	return nil
}
