package user

import "github.com/oklog/ulid/v2"

type Repository interface {
	GetById(id ulid.ULID) (*User, error)
	GetByEmailAddress(emailAddress string) (*User, error)
	GetAll(schoolId ulid.ULID) (*[]User, error)
	Add(user *User) error
	Remove(user *User) error
	Save(user *User) error
}
