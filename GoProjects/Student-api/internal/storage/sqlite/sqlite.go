package sqlite

// install go sqlite db driver "go get github.com/mattn/go-sqlite3"

import (
	"database/sql"
	"fmt"

	"github.com/Gauravji012/GoLang/GoProjects/Student-api/internal/config"
	"github.com/Gauravji012/GoLang/GoProjects/Student-api/internal/types"
	_ "github.com/mattn/go-sqlite3" // _ means we are not directly use this package
)

type Sqlite struct {
	Db *sql.DB
}

// creating instance of Sqlite struct because constructor method directly is not present in go

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}
	// create table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER
	)`)

	if err != nil {
		return nil, err // because sqlite instance is not created yet
	}

	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	/* some steps to create the record in db
	1. prepare the query
	*/

	stmt, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email, age) // execute the query & order matters
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}


func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id, name, email, age FROM students WHERE id = ? LIMIT 1")
	if err != nil {
		return types.Student{}, err
	}
	defer stmt.Close()
	// first we serialize/handle  the data in serialize manner that comes in stmt variable after executing the query

	var student types.Student
	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student. Age)
	if err != nil {
		if err == sql.ErrNoRows{
			return types.Student{}, fmt.Errorf("no student found with id %s", fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("query error: %w", err)
	}
	return student, nil

}