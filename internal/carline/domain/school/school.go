package school

import (
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/domain/student"
	_type "github.com/pascalallen/carline/internal/carline/infrastructure/database/type"
	"time"
)

type School struct {
	Id         _type.GormUlid    `json:"id" gorm:"primaryKey;size:26;not null"`
	Name       string            `json:"name" gorm:"size:100;not null"`
	Students   []student.Student `json:"students" gorm:"foreignKey:SchoolId;references:Id"`
	CreatedAt  time.Time         `json:"created_at" gorm:"not null"`
	ModifiedAt *time.Time        `json:"modified_at,omitempty" gorm:"default:null"`
	DeletedAt  *time.Time        `json:"deleted_at,omitempty" gorm:"default:null"`
}

func Create(id ulid.ULID, name string) *School {
	createdAt := time.Now()

	return &School{
		Id:        _type.GormUlid(id),
		Name:      name,
		CreatedAt: createdAt,
	}
}
