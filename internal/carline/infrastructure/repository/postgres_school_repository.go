package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/domain/school"
)

type PostgresSchoolRepository struct {
	session *sql.DB
}

func NewPostgresSchoolRepository(session *sql.DB) school.Repository {
	return &PostgresSchoolRepository{
		session: session,
	}
}

func (r *PostgresSchoolRepository) GetById(id ulid.ULID) (*school.School, error) {
	var s school.School
	var i string

	row := r.session.QueryRow("SELECT * FROM schools WHERE id = $1", id.String())
	if err := row.Scan(&i, &s.Name, &s.CreatedAt, &s.ModifiedAt, &s.DeletedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning School by ID: %s", err)
	}

	s.Id = ulid.MustParse(i)

	// TODO: eager load students

	return &s, nil
}

func (r *PostgresSchoolRepository) GetByName(name string) (*school.School, error) {
	var s school.School
	var id string

	row := r.session.QueryRow("SELECT * FROM schools WHERE name = $1", name)
	if err := row.Scan(&id, &s.Name, &s.CreatedAt, &s.ModifiedAt, &s.DeletedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning School by name: %s", err)
	}

	s.Id = ulid.MustParse(id)

	// TODO: eager load students

	return &s, nil
}

// GetAll TODO: Add pagination
func (r *PostgresSchoolRepository) GetAll(includeDeleted bool) (*[]school.School, error) {
	var s school.School
	var schools []school.School
	var id string
	query := "SELECT * FROM schools"

	if !includeDeleted {
		query += " WHERE deleted_at IS NULL"
	}

	rows, err := r.session.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching all Schools: %s", err)
	}

	for rows.Next() {
		if err := rows.Scan(&id, &s.Name, &s.CreatedAt, &s.ModifiedAt, &s.DeletedAt); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}

			return nil, fmt.Errorf("error scanning all Schools: %s", err)
		}

		s.Id = ulid.MustParse(id)
		schools = append(schools, s)
	}

	return &schools, nil
}

func (r *PostgresSchoolRepository) Add(school *school.School) error {
	if _, err := r.session.Exec("INSERT INTO schools(id, name, created_at) VALUES($1, $2, $3)", school.Id.String(), school.Name, school.CreatedAt); err != nil {
		return fmt.Errorf("failed to persist School to database: %v", err)
	}

	return nil
}

func (r *PostgresSchoolRepository) Remove(school *school.School) error {
	school.Delete()

	if _, err := r.session.Exec("UPDATE schools SET deleted_at = $1 WHERE id = $2", school.DeletedAt, school.Id); err != nil {
		return fmt.Errorf("failed to soft delete School: %v", err)
	}

	return nil
}
