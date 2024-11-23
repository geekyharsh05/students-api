package sqlite

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/geekyharsh05/students-api/internal/config"
	"github.com/geekyharsh05/students-api/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		age INTEGER NOT NULL
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	slog.Info("Student created successfully", slog.Int64("userId", lastId))

	return lastId, nil
}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT * FROM students WHERE id = ? LIMIT 1")
	if err != nil {
		return types.Student{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var student types.Student

	err = row.Scan(&student.Id, &student.Name, &student.Email, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("student with id %d not found", id)
		}

		return types.Student{}, fmt.Errorf("failed to scan row: %w", err)
	}

	return student, nil
}
