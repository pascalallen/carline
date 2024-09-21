package command

import (
	"github.com/oklog/ulid/v2"
)

type AddSchool struct {
	Id   ulid.ULID `json:"id"`
	Name string    `json:"name"`
}

func (c AddSchool) CommandName() string {
	return "AddSchool"
}
