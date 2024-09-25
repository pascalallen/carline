package command

import (
	"github.com/oklog/ulid/v2"
)

type ImportStudents struct {
	SchoolId   ulid.ULID `json:"school_id"`
	FileBuffer []byte    `json:"file_buffer"`
}

func (c ImportStudents) CommandName() string {
	return "ImportStudents"
}

type DeleteStudent struct {
	SchoolId ulid.ULID `json:"school_id"`
	Id       ulid.ULID `json:"id"`
}

func (c DeleteStudent) CommandName() string {
	return "DeleteStudent"
}
