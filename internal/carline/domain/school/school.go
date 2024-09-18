package school

import (
	"github.com/oklog/ulid/v2"
	_type "github.com/pascalallen/carline/internal/carline/infrastructure/database/type"
	"time"
)

type School struct {
	Id         _type.GormUlid `json:"id" gorm:"primaryKey;size:26;not null"`
	Name       string         `json:"name" gorm:"size:100;not null"`
	CreatedAt  time.Time      `json:"created_at" gorm:"not null"`
	ModifiedAt *time.Time     `json:"modified_at,omitempty" gorm:"default:null"`
	DeletedAt  *time.Time     `json:"deleted_at,omitempty" gorm:"default:null"` // TODO: Make nullable/optional
}

type SchoolRepository interface {
	GetById(id ulid.ULID) (*School, error)
	GetByName(name string) (*School, error)
	GetAll(includeDeleted bool) (*[]School, error)
	Add(school *School) error
	Remove(school *School) error
	UpdateOrAdd(school *School) error
}
