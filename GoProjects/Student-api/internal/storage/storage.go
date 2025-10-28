package storage

import "github.com/Gauravji012/GoLang/GoProjects/Student-api/internal/types"

// interfaces is an important concept. with the help of interfaces we can make plugins easily. we can handle failures , switching easily.
// we can switch from one application to another one easily. like sqlite -> postgresql

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
}
