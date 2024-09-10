package command

import "github.com/oklog/ulid/v2"

type UpdateUserEmailAddress struct {
	Id           ulid.ULID `json:"id"`
	EmailAddress string    `json:"email_address"`
}

func (c UpdateUserEmailAddress) CommandName() string {
	return "UpdateUserEmailAddress"
}
