package command

import (
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/domain/crypto"
)

type SendWelcomeEmail struct {
	Id           ulid.ULID     `json:"id"`
	FirstName    string        `json:"first_name"`
	LastName     string        `json:"last_name"`
	EmailAddress string        `json:"email_address"`
	Token        crypto.Crypto `json:"token"`
}

func (c SendWelcomeEmail) CommandName() string {
	return "SendWelcomeEmail"
}
