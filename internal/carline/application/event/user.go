package event

import (
	"github.com/oklog/ulid/v2"
)

type UserRegistered struct {
	Id              ulid.ULID `json:"id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	EmailAddress    string    `json:"email_address"`
	SecurityTokenId ulid.ULID `json:"security_token_id"`
}

func (e UserRegistered) EventName() string {
	return "UserRegistered"
}

type UserUpdated struct {
	Id           ulid.ULID `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	EmailAddress string    `json:"email_address"`
}

func (e UserUpdated) EventName() string {
	return "UserUpdated"
}
