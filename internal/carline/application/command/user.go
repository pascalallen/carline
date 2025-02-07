package command

import (
	"github.com/oklog/ulid/v2"
)

type RegisterUser struct {
	Id           ulid.ULID `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	EmailAddress string    `json:"email_address"`
	Admin        bool      `json:"admin"`
}

func (c RegisterUser) CommandName() string {
	return "RegisterUser"
}

type UpdateUser struct {
	Id           ulid.ULID
	FirstName    string
	LastName     string
	EmailAddress string
}

func (c UpdateUser) CommandName() string {
	return "UpdateUser"
}
