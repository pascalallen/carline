package student

import (
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/domain/school"
	"time"
)

type Student struct {
	Id         ulid.ULID     `json:"id"`
	TagNumber  string        `json:"tag_number"`
	FirstName  string        `json:"first_name"`
	LastName   string        `json:"last_name"`
	School     school.School `json:"school"`
	CreatedAt  time.Time     `json:"created_at"`
	ModifiedAt *time.Time    `json:"modified_at,omitempty"`
	DeletedAt  *time.Time    `json:"deleted_at,omitempty"`
}

func Register(id ulid.ULID, tagNumber string, firstName string, lastName string) *Student {
	createdAt := time.Now()

	return &Student{
		Id:        id,
		TagNumber: tagNumber,
		FirstName: firstName,
		LastName:  lastName,
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
	return !s.DeletedAt.IsZero()
}

func (s *Student) Delete() {
	now := time.Now()
	s.DeletedAt = &now
	s.ModifiedAt = s.DeletedAt
}

func (s *Student) Restore() {
	now := time.Now()
	s.DeletedAt = nil
	s.ModifiedAt = &now
}
