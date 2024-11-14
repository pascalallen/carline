package query

import "github.com/oklog/ulid/v2"

type GetSchoolByName struct {
	Name string `json:"name"`
}

func (q GetSchoolByName) QueryName() string {
	return "GetSchoolByName"
}

type ListSchools struct {
	UserId ulid.ULID `json:"user_id"`
}

func (q ListSchools) QueryName() string {
	return "ListSchools"
}

type GetSchoolByIdAndUserId struct {
	UserId ulid.ULID `json:"user_id"`
	Id     ulid.ULID `json:"id"`
}

func (q GetSchoolByIdAndUserId) QueryName() string {
	return "GetSchoolByIdAndUserId"
}
