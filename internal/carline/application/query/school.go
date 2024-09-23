package query

import "github.com/oklog/ulid/v2"

type GetSchoolById struct {
	Id ulid.ULID `json:"id"`
}

func (q GetSchoolById) QueryName() string {
	return "GetSchoolById"
}

type GetSchoolByName struct {
	Name string `json:"name"`
}

func (q GetSchoolByName) QueryName() string {
	return "GetSchoolByName"
}

type ListSchools struct {
	IncludeDeleted bool `json:"include_deleted"`
}

func (q ListSchools) QueryName() string {
	return "ListSchools"
}
