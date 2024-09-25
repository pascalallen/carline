package query

import "github.com/oklog/ulid/v2"

type ListStudents struct {
	SchoolId       ulid.ULID `json:"school_id"`
	IncludeDeleted bool      `json:"include_deleted"`
}

func (q ListStudents) QueryName() string {
	return "ListStudents"
}

type GetStudentById struct {
	SchoolId ulid.ULID `json:"school_id"`
	Id       ulid.ULID `json:"id"`
}

func (q GetStudentById) QueryName() string {
	return "GetStudentById"
}
