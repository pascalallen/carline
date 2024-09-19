package repository

import (
	"errors"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/domain/school"
	"github.com/pascalallen/carline/internal/carline/infrastructure/database"
	"gorm.io/gorm"
)

type GormSchoolRepository struct {
	session database.Session
}

func NewGormSchoolRepository(session database.Session) school.SchoolRepository {
	return &GormSchoolRepository{
		session: session,
	}
}

func (r *GormSchoolRepository) GetById(id ulid.ULID) (*school.School, error) {
	var s *school.School
	if err := r.session.First(&s, "id = ?", id.String()); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to fetch School by ID: %s", id)
	}

	return s, nil
}

func (r *GormSchoolRepository) GetByName(name string) (*school.School, error) {
	var s *school.School
	if err := r.session.First(&s, "name = ?", name); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to fetch School by name: %s", name)
	}

	return s, nil
}

// GetAll TODO: Add pagination
func (r *GormSchoolRepository) GetAll(includeDeleted bool) (*[]school.School, error) {
	var schools *[]school.School
	if !includeDeleted {
		r.session = r.session.Where("deleted_at IS NULL")
	}

	if err := r.session.Find(&schools); err != nil {
		return nil, fmt.Errorf("failed to fetch all Schools: %s", err)
	}

	return schools, nil
}

func (r *GormSchoolRepository) Add(school *school.School) error {
	if err := r.session.Create(school); err != nil {
		return fmt.Errorf("failed to persist School to database: %s", school)
	}

	return nil
}

func (r *GormSchoolRepository) Remove(school *school.School) error {
	if err := r.session.Delete(&school); err != nil {
		return fmt.Errorf("failed to delete School from database: %s", school)
	}

	return nil
}

func (r *GormSchoolRepository) UpdateOrAdd(school *school.School) error {
	if err := r.session.Save(&school); err != nil {
		return fmt.Errorf("failed to update School: %s", school)
	}

	return nil
}
