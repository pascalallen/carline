package student

import "github.com/oklog/ulid/v2"

type Repository interface {
	GetById(id ulid.ULID) (*Student, error)
	GetByTagNumber(tagNumber string) (*Student, error)
	GetAllBySchoolIdAndTagNumber(schoolId ulid.ULID, tagNumber string) (*[]Student, error)
	GetAll(schoolId ulid.ULID, dismissed bool) (*[]Student, error)
	Add(student *Student) error
	Remove(student *Student) error
	Dismiss(student *Student) error
}
