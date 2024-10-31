package command_handler

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/application/command"
	"github.com/pascalallen/carline/internal/carline/domain/school"
	"github.com/pascalallen/carline/internal/carline/domain/student"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ImportStudentsHandler struct {
	SchoolRepository  school.Repository
	StudentRepository student.Repository
	DatabaseSession   *sql.DB
}

var requiredHeaders = []string{"tag_number", "first_name", "last_name"}

func (h ImportStudentsHandler) Handle(cmd messaging.Command) error {
	c, ok := cmd.(*command.ImportStudents)
	if !ok {
		return fmt.Errorf("invalid command type passed to ImportStudentsHandler: %v", cmd)
	}

	s, err := h.SchoolRepository.GetById(c.SchoolId)
	if err != nil {
		return fmt.Errorf("school ID not found: %s", s)
	}

	// Create a temp file to store the uploaded file
	destPath := filepath.Join("/tmp", fmt.Sprintf("%s-students.csv", c.SchoolId))
	dstFile, err := os.Create(destPath)
	if err != nil {
		return errors.New("failed to create destination file")
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, strings.NewReader(string(c.FileBuffer))); err != nil {
		return errors.New("failed to copy file to temp filesystem")
	}

	if _, err := dstFile.Seek(0, 0); err != nil {
		return errors.New("failed to reset file pointer for parsing")
	}

	// Reopen the file for parsing
	file, err := os.Open(destPath)
	if err != nil {
		return errors.New("failed to open temp file for reading")
	}
	defer file.Close()

	reader := csv.NewReader(file)

	headers, err := reader.Read()
	if err != nil {
		return errors.New("failed to read CSV headers")
	}

	missingHeaders := checkHeaders(headers, requiredHeaders)
	if len(missingHeaders) > 0 {
		return fmt.Errorf("missing required CSV headers: %v", missingHeaders)
	}

	headerMap := make(map[string]int)
	for idx, header := range headers {
		headerMap[header] = idx
	}

	rows, err := reader.ReadAll()
	if err != nil {
		return errors.New("failed to read CSV body")
	}

	tx, err := h.DatabaseSession.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %s", err)
	}

	stmt, err := tx.Prepare("INSERT INTO students(id, tag_number, first_name, last_name, school_id, created_at) VALUES($1, $2, $3, $4, $5, $6)")
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	for _, row := range rows {
		id := ulid.Make()
		_, err := stmt.Exec(id.String(), row[headerMap["tag_number"]], row[headerMap["first_name"]], row[headerMap["last_name"]], s.Id.String(), time.Now())
		if err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("failed to execute statement: %s", err)
		}
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("failed to commit transaction: %s", err)
	}

	return nil
}

func checkHeaders(headers []string, requiredHeaders []string) []string {
	headerMap := make(map[string]bool)
	for _, header := range headers {
		headerMap[header] = true
	}

	var missingHeaders []string
	for _, required := range requiredHeaders {
		if !headerMap[required] {
			missingHeaders = append(missingHeaders, required)
		}
	}

	return missingHeaders
}

type DeleteStudentHandler struct {
	StudentRepository student.Repository
}

func (h DeleteStudentHandler) Handle(cmd messaging.Command) error {
	c, ok := cmd.(*command.DeleteStudent)
	if !ok {
		return fmt.Errorf("invalid command type passed to DeleteStudentHandler: %v", cmd)
	}

	s, err := h.StudentRepository.GetById(c.Id)
	if err != nil {
		return fmt.Errorf("student not found: %s", c.Id)
	}

	err = h.StudentRepository.Remove(s)
	if err != nil {
		return fmt.Errorf("student removal failed: %s", err)
	}

	return nil
}
