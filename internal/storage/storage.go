package storage

import "github.com/AnshKumar10/Go-Student-Management-Api/internal/types"

type Storage interface {
	GetAllStudents() ([]types.Student, error)
	GetStudentById(id int64) (types.Student, error)
	CreateStudent(name string, email string, age int) (int64, error)
	UpdateStudentById(id int64, name string, email string, age int) (int64, error)
	DeleteStudentById(id int64) (bool, error)
}
