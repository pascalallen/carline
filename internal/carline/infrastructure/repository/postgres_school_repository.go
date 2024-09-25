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
	q := `SELECT 
			id,
			name,
			created_at,
			modified_at,
			deleted_at
		FROM schools 
		WHERE id = $1;`

	row := r.session.QueryRow(q, id.String())
	if err := row.Scan(&i, &s.Name, &s.CreatedAt, &s.ModifiedAt, &s.DeletedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning School by ID: %s", err)
	}

	s.Id = ulid.MustParse(i)

	return &s, nil
}

func (r *PostgresSchoolRepository) GetByName(name string) (*school.School, error) {
	var s school.School
	var id string
	q := `SELECT 
			id,
			name,
			created_at,
			modified_at,
			deleted_at
		FROM schools 
		WHERE name = $1;`

	row := r.session.QueryRow(q, name)
	if err := row.Scan(&id, &s.Name, &s.CreatedAt, &s.ModifiedAt, &s.DeletedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning School by name: %s", err)
	}

	s.Id = ulid.MustParse(id)

	return &s, nil
}

func (r *PostgresSchoolRepository) GetAll(includeDeleted bool) (*[]school.School, error) {
	var schools []school.School
	q := `SELECT 
			id,
			name,
			created_at,
			modified_at,
			deleted_at
		FROM schools`

	if !includeDeleted {
		q += ` WHERE deleted_at IS NULL;`
	}

	rows, err := r.session.Query(q)
	if err != nil {
		return nil, fmt.Errorf("error fetching all Schools: %s", err)
	}

	for rows.Next() {
		var id string
		var s school.School

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
	q := `INSERT INTO schools(id, name, created_at) VALUES($1, $2, $3)`

	if _, err := r.session.Exec(q, school.Id.String(), school.Name, school.CreatedAt); err != nil {
		return fmt.Errorf("failed to persist School to database: %v", err)
	}

	return nil
}

func (r *PostgresSchoolRepository) Remove(school *school.School) error {
	school.Delete()
	q := `UPDATE schools SET deleted_at = $1, modified_at = $2 WHERE id = $3`

	if _, err := r.session.Exec(q, school.DeletedAt, school.ModifiedAt, school.Id.String()); err != nil {
		return fmt.Errorf("failed to soft delete School: %v", err)
	}

	return nil
}
