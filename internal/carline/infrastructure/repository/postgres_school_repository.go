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
			modified_at
		FROM schools 
		WHERE id = $1;`

	row := r.session.QueryRow(q, id.String())
	if err := row.Scan(&i, &s.Name, &s.CreatedAt, &s.ModifiedAt); err != nil {
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
			modified_at
		FROM schools 
		WHERE name = $1;`

	row := r.session.QueryRow(q, name)
	if err := row.Scan(&id, &s.Name, &s.CreatedAt, &s.ModifiedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning School by name: %s", err)
	}

	s.Id = ulid.MustParse(id)

	return &s, nil
}

func (r *PostgresSchoolRepository) GetAll(userId ulid.ULID) (*[]school.School, error) {
	var schools []school.School

	q := `SELECT 
			id,
			name,
			created_at,
			modified_at
		FROM schools
		JOIN user_schools ON user_schools.school_id = schools.id
		WHERE user_schools.user_id = $1;`

	rows, err := r.session.Query(q, userId.String())
	if err != nil {
		return nil, fmt.Errorf("error fetching all Schools: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var s school.School

		if err := rows.Scan(&id, &s.Name, &s.CreatedAt, &s.ModifiedAt); err != nil {
			return nil, fmt.Errorf("error scanning School record: %v", err)
		}

		s.Id = ulid.MustParse(id)
		schools = append(schools, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return &schools, nil
}

func (r *PostgresSchoolRepository) Add(school *school.School, userId ulid.ULID) error {
	tx, err := r.session.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction for PostgresSchoolRepository.Add(): %v", err)
	}

	q := `INSERT INTO schools(id, name, created_at) VALUES($1, $2, $3)`
	_, err = tx.Exec(q, school.Id.String(), school.Name, school.CreatedAt)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("failed to persist School to database: %v", err)
	}

	q = `INSERT INTO user_schools(user_id, school_id) VALUES($1, $2)`
	_, err = tx.Exec(q, userId.String(), school.Id.String())
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("failed to persist user_schools record to database: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction for PostgresSchoolRepository.Add(): %v", err)
	}

	return nil
}

func (r *PostgresSchoolRepository) Remove(school *school.School, userId ulid.ULID) error {
	tx, err := r.session.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction for PostgresSchoolRepository.Remove(): %v", err)
	}

	q := `DELETE FROM user_schools WHERE user_id = $1 AND school_id = $2`
	_, err = tx.Exec(q, userId.String(), school.Id.String())
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("failed to remove user_schools record from database: %v", err)
	}

	q = `DELETE FROM schools WHERE id = $1`
	_, err = tx.Exec(q, school.Id.String())
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("failed to remove School from database: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction for PostgresSchoolRepository.Remove(): %v", err)
	}

	return nil
}
