package school

import (
	"github.com/oklog/ulid/v2"
	"time"
)

type School struct {
	Id         ulid.ULID  `json:"id"`
	Name       string     `json:"name"`
	CreatedAt  time.Time  `json:"created_at"`
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

func Create(id ulid.ULID, name string) *School {
	createdAt := time.Now()

	return &School{
		Id:        id,
		Name:      name,
		CreatedAt: createdAt,
	}
}

func (s *School) UpdateName(name string) {
	s.Name = name
	now := time.Now()
	s.ModifiedAt = &now
}

func (s *School) IsDeleted() bool {
	return s.DeletedAt != nil
}

func (s *School) Delete() {
	now := time.Now()
	s.DeletedAt = &now
	s.ModifiedAt = &now
}

func (s *School) Restore() {
	now := time.Now()
	s.DeletedAt = nil
	s.ModifiedAt = &now
}
