package command

import (
	"github.com/oklog/ulid/v2"
)

type SendWelcomeEmail struct {
	Id              ulid.ULID `json:"id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	EmailAddress    string    `json:"email_address"`
	SecurityTokenId ulid.ULID `json:"security_token_id"`
}

func (c SendWelcomeEmail) CommandName() string {
	return "SendWelcomeEmail"
}
