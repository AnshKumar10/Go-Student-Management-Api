package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/AnshKumar10/Go-Student-Management-Api/internal/config"
	"github.com/AnshKumar10/Go-Student-Management-Api/internal/types"
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
    name TEXT,
    email TEXT,
    age INTEGER
)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{Db: db}, nil

}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	statement, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?, ?, ?)")

	if err != nil {
		return 0, nil
	}

	defer statement.Close()

	result, err := statement.Exec(name, email, age)

	if err != nil {
		return 0, nil
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, nil
	}

	return lastId, nil
}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {

	statement, err := s.Db.Prepare("SELECT * FROM students WHERE id = ? LIMIT 1")

	if err != nil {
		return types.Student{}, nil
	}

	defer statement.Close()

	var student types.Student

	err = statement.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with id %s", fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("query Error %w", err)
	}

	return student, nil
}

func (s *Sqlite) UpdateStudentById(id int64, name string, email string, age int) (int64, error) {

	statement, err := s.Db.Prepare("UPDATE students SET name = ?, email = ?, age = ? WHERE id = ?")

	if err != nil {
		return 0, fmt.Errorf("prepare error: %w", err)
	}

	defer statement.Close()

	result, err := statement.Exec(name, email, age, id)

	if err != nil {
		return 0, fmt.Errorf("execution error: %w", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, fmt.Errorf("rows affected error: %w", err)
	}

	if rowsAffected == 0 {
		return 0, fmt.Errorf("no student found with id %d", id)
	}

	return rowsAffected, nil
}

func (s *Sqlite) DeleteStudentById(id int64) (bool, error) {

	statement, err := s.Db.Prepare("DELETE FROM students WHERE id = ?")

	if err != nil {
		return false, fmt.Errorf("prepare error: %w", err)
	}

	defer statement.Close()

	result, err := statement.Exec(id)

	if err != nil {
		return false, fmt.Errorf("execution error: %w", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return false, fmt.Errorf("rows affected error: %w", err)
	}

	if rowsAffected == 0 {
		return false, fmt.Errorf("no student found with id %d", id)
	}

	return true, nil
}

func (s *Sqlite) GetAllStudents() ([]types.Student, error) {

	statement, err := s.Db.Prepare("SELECT * FROM students")

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.Query()

	if err != nil {
		return nil, err
	}

	var students []types.Student

	for rows.Next() {
		var student types.Student

		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)

		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}
	return students, nil
}
