package command

import "github.com/oklog/ulid/v2"

type CreateSchool struct {
	Id     ulid.ULID `json:"id"`
	Name   string    `json:"name"`
	UserId ulid.ULID `json:"user_id"`
}

func (c CreateSchool) CommandName() string {
	return "CreateSchool"
}

type DeleteSchool struct {
	Id     ulid.ULID `json:"id"`
	UserId ulid.ULID `json:"user_id"`
}

func (c DeleteSchool) CommandName() string {
	return "DeleteSchool"
}
