package user

import "github.com/oklog/ulid/v2"

type Repository interface {
	GetById(id ulid.ULID) (*User, error)
	GetByEmailAddress(emailAddress string) (*User, error)
	GetAll(includeDeleted bool) (*[]User, error)
	Add(user *User) error
	Remove(user *User) error
}
