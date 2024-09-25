package command

import "github.com/oklog/ulid/v2"

type CreateSchool struct {
	Id   ulid.ULID `json:"id"`
	Name string    `json:"name"`
}

func (c CreateSchool) CommandName() string {
	return "CreateSchool"
}

type DeleteSchool struct {
	Id ulid.ULID `json:"id"`
}

func (c DeleteSchool) CommandName() string {
	return "DeleteSchool"
}
