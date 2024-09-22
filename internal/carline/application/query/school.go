package query

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
