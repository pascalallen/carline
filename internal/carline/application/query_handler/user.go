package query_handler

import (
	"fmt"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/domain/user"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/storage"
)

type ProjectionState struct {
	EmailAddresses map[string]string `json:"email_addresses"`
}

type GetUserByIdHandler struct {
	EventStore storage.EventStore
}

func (h GetUserByIdHandler) Handle(qry messaging.Query) (any, error) {
	q, ok := qry.(query.GetUserById)
	if !ok {
		return nil, fmt.Errorf("invalid query type passed to GetUserByIdHandler: %v", qry)
	}

	streamId := fmt.Sprintf("user-%s", q.Id)
	events, err := h.EventStore.ReadFromStream(streamId)
	if err != nil {
		return nil, fmt.Errorf("error attempting to read events from stream: %s", err)
	}

	if len(events) == 0 {
		return nil, fmt.Errorf("no events found for user ID: %s", q.Id)
	}

	u := user.LoadUserFromEvents(events)

	return u, nil
}

type GetUserByEmailAddressHandler struct {
	EventStore storage.EventStore
}

func (h GetUserByEmailAddressHandler) Handle(qry messaging.Query) (any, error) {
	q, ok := qry.(query.GetUserByEmailAddress)
	if !ok {
		return nil, fmt.Errorf("invalid query type passed to GetUserByEmailAddressHandler: %v", qry)
	}

	var result ProjectionState
	if err := h.EventStore.UnmarshalProjectionResult("user-email-addresses", &result); err != nil {
		return nil, fmt.Errorf("error getting projection result: %v", err)
	}

	userId, exists := result.EmailAddresses[q.EmailAddress]
	if !exists {
		return nil, fmt.Errorf("no user found with email address: %s", q.EmailAddress)
	}

	streamId := fmt.Sprintf("user-%s", userId)
	events, err := h.EventStore.ReadFromStream(streamId)
	if err != nil {
		return nil, fmt.Errorf("error attempting to read events from stream: %s", err)
	}

	if len(events) == 0 {
		return nil, fmt.Errorf("no events found for user ID: %s", userId)
	}

	u := user.LoadUserFromEvents(events)

	return u, nil
}
