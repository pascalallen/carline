package student

import "github.com/oklog/ulid/v2"

type Repository interface {
	GetById(id ulid.ULID) (*Student, error)
	GetByTagNumber(tagNumber string) (*Student, error)
	GetAll(includeDeleted bool) (*[]Student, error)
	Add(student *Student) error
	Remove(student *Student) error
	UpdateOrAdd(student *Student) error
}
