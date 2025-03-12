package command

import (
	"github.com/oklog/ulid/v2"
)

type RegisterUser struct {
	Id           ulid.ULID  `json:"id"`
	SchoolId     *ulid.ULID `json:"school_id,omitempty"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	EmailAddress string     `json:"email_address"`
	Admin        bool       `json:"admin"`
}

func (c RegisterUser) CommandName() string {
	return "RegisterUser"
}

type UpdateUser struct {
	Id           ulid.ULID `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	EmailAddress string    `json:"email_address"`
}

func (c UpdateUser) CommandName() string {
	return "UpdateUser"
}
