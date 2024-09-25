package student

import (
	"github.com/oklog/ulid/v2"
	"time"
)

type Student struct {
	Id         ulid.ULID  `json:"id"`
	TagNumber  string     `json:"tag_number"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	SchoolId   ulid.ULID  `json:"school_id"`
	CreatedAt  time.Time  `json:"created_at"`
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

func Register(id ulid.ULID, tagNumber string, firstName string, lastName string, schoolId ulid.ULID) *Student {
	createdAt := time.Now()

	return &Student{
		Id:        id,
		TagNumber: tagNumber,
		FirstName: firstName,
		LastName:  lastName,
		SchoolId:  schoolId,
		CreatedAt: createdAt,
	}
}

func (s *Student) UpdateTagNumber(tagNumber string) {
	s.TagNumber = tagNumber
	now := time.Now()
	s.ModifiedAt = &now
}

func (s *Student) UpdateFirstName(firstName string) {
	s.FirstName = firstName
	now := time.Now()
	s.ModifiedAt = &now
}

func (s *Student) UpdateLastName(lastName string) {
	s.LastName = lastName
	now := time.Now()
	s.ModifiedAt = &now
}

func (s *Student) IsDeleted() bool {
	return s.DeletedAt != nil
}

func (s *Student) Delete() {
	now := time.Now()
	s.DeletedAt = &now
	s.ModifiedAt = &now
}

func (s *Student) Restore() {
	now := time.Now()
	s.DeletedAt = nil
	s.ModifiedAt = &now
}
