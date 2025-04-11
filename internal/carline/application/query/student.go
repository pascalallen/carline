package query

import "github.com/oklog/ulid/v2"

type ListStudents struct {
	SchoolId  ulid.ULID `json:"school_id"`
	Dismissed bool      `json:"dismissed"`
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
