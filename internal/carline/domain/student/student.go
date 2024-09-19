package student

import (
	"github.com/oklog/ulid/v2"
	_type "github.com/pascalallen/carline/internal/carline/infrastructure/database/type"
	"time"
)

type Student struct {
	Id         _type.GormUlid `json:"id" gorm:"primaryKey;size:26;not null"`
	TagNumber  string         `json:"tag_number" gorm:"size:100;not null"`
	FirstName  string         `json:"first_name" gorm:"size:100;not null"`
	LastName   string         `json:"last_name" gorm:"size:100;not null"`
	CreatedAt  time.Time      `json:"created_at" gorm:"not null"`
	ModifiedAt *time.Time     `json:"modified_at,omitempty" gorm:"default:null"`
	DeletedAt  *time.Time     `json:"deleted_at,omitempty" gorm:"default:null"`
}

type StudentRepository interface {
	GetById(id ulid.ULID) (*Student, error)
	GetByTagNumber(tagNumber string) (*Student, error)
	GetAll(includeDeleted bool) (*[]Student, error)
	Add(student *Student) error
	Remove(student *Student) error
	UpdateOrAdd(student *Student) error
}
