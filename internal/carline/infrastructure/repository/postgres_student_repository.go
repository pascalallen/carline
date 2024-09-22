package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/domain/student"
)

type PostgresStudentRepository struct {
	session *sql.DB
}

func NewPostgresStudentRepository(session *sql.DB) student.Repository {
	return &PostgresStudentRepository{
		session: session,
	}
}

func (r *PostgresStudentRepository) GetById(id ulid.ULID) (*student.Student, error) {
	var s student.Student
	var i string

	row := r.session.QueryRow("SELECT * FROM students WHERE id = $1", id.String())
	if err := row.Scan(&i, &s.TagNumber, &s.FirstName, &s.LastName, &s.CreatedAt, &s.ModifiedAt, &s.DeletedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning Student by ID: %s", err)
	}

	s.Id = ulid.MustParse(i)

	// TODO: eager load school

	return &s, nil
}

func (r *PostgresStudentRepository) GetByTagNumber(tagNumber string) (*student.Student, error) {
	var s student.Student
	var id string

	row := r.session.QueryRow("SELECT * FROM students WHERE tag_number = $1", tagNumber)
	if err := row.Scan(&id, &s.TagNumber, &s.FirstName, &s.LastName, &s.CreatedAt, &s.ModifiedAt, &s.DeletedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning Student by tag number: %s", err)
	}

	s.Id = ulid.MustParse(id)

	// TODO: eager load school

	return &s, nil
}

// GetAll TODO: Add pagination
func (r *PostgresStudentRepository) GetAll(includeDeleted bool) (*[]student.Student, error) {
	var s student.Student
	var students []student.Student
	var id string
	query := "SELECT * FROM students"

	if !includeDeleted {
		query += " WHERE deleted_at IS NULL"
	}

	rows, err := r.session.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching all Students: %s", err)
	}

	for rows.Next() {
		if err := rows.Scan(&id, &s.TagNumber, &s.FirstName, &s.LastName, &s.CreatedAt, &s.ModifiedAt, &s.DeletedAt); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}

			return nil, fmt.Errorf("error scanning all Students: %s", err)
		}

		s.Id = ulid.MustParse(id)
		students = append(students, s)
	}

	return &students, nil
}

func (r *PostgresStudentRepository) Add(student *student.Student) error {
	if _, err := r.session.Exec("INSERT INTO students(id, tag_number, first_name, last_name, created_at) VALUES($1, $2, $3, $4, $5)", student.Id.String(), student.TagNumber, student.FirstName, student.LastName, student.CreatedAt); err != nil {
		return fmt.Errorf("failed to persist Student to database: %v", err)
	}

	return nil
}

func (r *PostgresStudentRepository) Remove(student *student.Student) error {
	student.Delete()

	if _, err := r.session.Exec("UPDATE students SET deleted_at = $1 WHERE id = $2", student.DeletedAt, student.Id); err != nil {
		return fmt.Errorf("failed to soft delete Student: %v", err)
	}

	return nil
}
