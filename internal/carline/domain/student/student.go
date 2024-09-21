package student

import (
	_type "github.com/pascalallen/carline/internal/carline/infrastructure/database/type"
	"time"
)

type Student struct {
	Id         _type.GormUlid `json:"id" gorm:"primaryKey;size:26;not null"`
	SchoolId   _type.GormUlid `json:"school_id" gorm:"index"`
	TagNumber  string         `json:"tag_number" gorm:"size:100;not null"`
	FirstName  string         `json:"first_name" gorm:"size:100;not null"`
	LastName   string         `json:"last_name" gorm:"size:100;not null"`
	CreatedAt  time.Time      `json:"created_at" gorm:"not null"`
	ModifiedAt *time.Time     `json:"modified_at,omitempty" gorm:"default:null"`
	DeletedAt  *time.Time     `json:"deleted_at,omitempty" gorm:"default:null"`
}
