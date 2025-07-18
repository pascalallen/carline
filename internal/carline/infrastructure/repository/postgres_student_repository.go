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
	var sid string
	q := `SELECT 
			id, 
			tag_number, 
			first_name, 
			last_name, 
			dismissed,
			school_id, 
			created_at, 
			modified_at
		FROM students 
		WHERE id = $1;`

	row := r.session.QueryRow(q, id.String())
	if err := row.Scan(&i, &s.TagNumber, &s.FirstName, &s.LastName, &s.Dismissed, &sid, &s.CreatedAt, &s.ModifiedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning Student by ID: %s", err)
	}

	s.Id = ulid.MustParse(i)
	s.SchoolId = ulid.MustParse(sid)

	return &s, nil
}

func (r *PostgresStudentRepository) GetByTagNumber(tagNumber string) (*student.Student, error) {
	var s student.Student
	var id string
	q := `SELECT 
			id, 
			tag_number, 
			first_name, 
			last_name, 
			dismissed,
			school_id, 
			created_at, 
			modified_at
		FROM students 
		WHERE tag_number = $1;`

	row := r.session.QueryRow(q, tagNumber)
	if err := row.Scan(&id, &s.TagNumber, &s.FirstName, &s.LastName, &s.Dismissed, &s.SchoolId, &s.CreatedAt, &s.ModifiedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning Student by tag number: %s", err)
	}

	s.Id = ulid.MustParse(id)

	return &s, nil
}

func (r *PostgresStudentRepository) GetAllBySchoolIdAndTagNumber(schoolId ulid.ULID, tagNumber string) (*[]student.Student, error) {
	var students []student.Student
	q := `SELECT 
			id, 
			tag_number, 
			first_name, 
			last_name, 
			dismissed,
			school_id, 
			created_at, 
			modified_at
		FROM students 
		WHERE school_id = $1
		AND tag_number = $2`

	rows, err := r.session.Query(q, schoolId.String(), tagNumber)
	if err != nil {
		return nil, fmt.Errorf("error fetching all Students by School ID and Tag Number: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var sid string
		var id string
		var s student.Student

		if err := rows.Scan(&id, &s.TagNumber, &s.FirstName, &s.LastName, &s.Dismissed, &sid, &s.CreatedAt, &s.ModifiedAt); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}

			return nil, fmt.Errorf("error scanning all Students by School ID and Tag Number: %s", err)
		}

		s.Id = ulid.MustParse(id)
		s.SchoolId = ulid.MustParse(sid)
		students = append(students, s)
	}

	return &students, nil
}

func (r *PostgresStudentRepository) GetAll(schoolId ulid.ULID, dismissed bool) (*[]student.Student, error) {
	var students []student.Student
	q := `SELECT 
			id, 
			tag_number, 
			first_name, 
			last_name, 
			dismissed,
			school_id, 
			created_at, 
			modified_at
		FROM students 
		WHERE school_id = $1
		AND dismissed = $2`

	rows, err := r.session.Query(q, schoolId.String(), dismissed)
	if err != nil {
		return nil, fmt.Errorf("error fetching all Students: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var sid string
		var id string
		var s student.Student

		if err := rows.Scan(&id, &s.TagNumber, &s.FirstName, &s.LastName, &s.Dismissed, &sid, &s.CreatedAt, &s.ModifiedAt); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}

			return nil, fmt.Errorf("error scanning all Students: %s", err)
		}

		s.Id = ulid.MustParse(id)
		s.SchoolId = ulid.MustParse(sid)
		students = append(students, s)
	}

	return &students, nil
}

func (r *PostgresStudentRepository) Add(student *student.Student) error {
	q := `INSERT INTO students(id, tag_number, first_name, last_name, dismissed, school_id, created_at) VALUES($1, $2, $3, $4, $5, $6, $7)`

	if _, err := r.session.Exec(q, student.Id.String(), student.TagNumber, student.FirstName, student.LastName, student.Dismissed, student.SchoolId, student.CreatedAt); err != nil {
		return fmt.Errorf("failed to persist Student to database: %v", err)
	}

	return nil
}

func (r *PostgresStudentRepository) Remove(student *student.Student) error {
	q := `DELETE FROM students WHERE id = $1`

	if _, err := r.session.Exec(q, student.Id.String()); err != nil {
		return fmt.Errorf("failed to remove Student from database: %v", err)
	}

	return nil
}

func (r *PostgresStudentRepository) Dismiss(student *student.Student) error {
	q := `UPDATE students SET dismissed = $1, modified_at = $2 WHERE id = $3 AND school_id = $4 AND tag_number = $5`

	_, err := r.session.Exec(
		q,
		true,
		student.ModifiedAt,
		student.Id.String(),
		student.SchoolId.String(),
		student.TagNumber,
	)

	if err != nil {
		return fmt.Errorf("failed to dismiss Student in database: %v", err)
	}

	return nil
}
