package event

import (
	"github.com/oklog/ulid/v2"
)

type WelcomeEmailSent struct {
	Id              ulid.ULID `json:"id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	EmailAddress    string    `json:"email_address"`
	SecurityTokenId ulid.ULID `json:"security_token_id"`
}

func (e WelcomeEmailSent) EventName() string {
	return "WelcomeEmailSent"
}
