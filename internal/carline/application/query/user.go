package query

import "github.com/oklog/ulid/v2"

type ListUsers struct {
	SchoolId ulid.ULID `json:"school_id"`
}

func (q ListUsers) QueryName() string {
	return "ListUsers"
}

type GetUserById struct {
	Id ulid.ULID `json:"id"`
}

func (q GetUserById) QueryName() string {
	return "GetUserById"
}

type GetUserByEmailAddress struct {
	EmailAddress string `json:"email_address"`
}

func (q GetUserByEmailAddress) QueryName() string {
	return "GetUserByEmailAddress"
}
