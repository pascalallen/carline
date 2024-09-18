package event

import (
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/domain/password"
)

type UserRegistered struct {
	Id           ulid.ULID     `json:"id"`
	FirstName    string        `json:"first_name"`
	LastName     string        `json:"last_name"`
	EmailAddress string        `json:"email_address"`
	PasswordHash password.Hash `json:"password_hash"`
}

func (e UserRegistered) EventName() string {
	return "UserRegistered"
}

type UserEmailAddressUpdated struct {
	Id           ulid.ULID `json:"id"`
	EmailAddress string    `json:"email_address"`
}

func (e UserEmailAddressUpdated) EventName() string {
	return "UserEmailAddressUpdated"
}
