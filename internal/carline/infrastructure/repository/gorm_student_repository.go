package repository

import (
	"errors"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/domain/student"
	"github.com/pascalallen/carline/internal/carline/infrastructure/database"
	"gorm.io/gorm"
)

type GormStudentRepository struct {
	session database.Session
}

func NewGormStudentRepository(session database.Session) student.Repository {
	return &GormStudentRepository{
		session: session,
	}
}

func (r *GormStudentRepository) GetById(id ulid.ULID) (*student.Student, error) {
	var s *student.Student
	if err := r.session.First(&s, "id = ?", id.String()); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to fetch Student by ID: %s", id)
	}

	return s, nil
}

func (r *GormStudentRepository) GetByTagNumber(tagNumber string) (*student.Student, error) {
	var s *student.Student
	if err := r.session.First(&s, "tag_number = ?", tagNumber); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to fetch Student by tag number: %s", tagNumber)
	}

	return s, nil
}

// GetAll TODO: Add pagination
func (r *GormStudentRepository) GetAll(includeDeleted bool) (*[]student.Student, error) {
	var students *[]student.Student
	if !includeDeleted {
		r.session = r.session.Where("deleted_at IS NULL")
	}

	if err := r.session.Find(&students); err != nil {
		return nil, fmt.Errorf("failed to fetch all Students: %s", err)
	}

	return students, nil
}

func (r *GormStudentRepository) Add(student *student.Student) error {
	if err := r.session.Create(student); err != nil {
		return fmt.Errorf("failed to persist Student to database: %s", student)
	}

	return nil
}

func (r *GormStudentRepository) Remove(student *student.Student) error {
	if err := r.session.Delete(&student); err != nil {
		return fmt.Errorf("failed to delete Student from database: %s", student)
	}

	return nil
}

func (r *GormStudentRepository) UpdateOrAdd(student *student.Student) error {
	if err := r.session.Save(&student); err != nil {
		return fmt.Errorf("failed to update Student: %s", student)
	}

	return nil
}
