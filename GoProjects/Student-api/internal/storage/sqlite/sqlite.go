package sqlite

// install go sqlite db driver "go get github.com/mattn/go-sqlite3"

import (
	"database/sql"
	"fmt"
	// "strings"

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
	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with id %s", fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("query error: %w", err)
	}
	return student, nil

}

func (s *Sqlite) DeleteStudentFromDB(id int64) error {
	stmt, err := s.Db.Prepare("DELETE FROM students WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sqlite) GetStudents() ([]types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id, name, email, age FROM students")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query() // rows is basically a list
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

func (s *Sqlite) UpdateStudent(id int64, payload types.Student) error {
	// update every field other wise update missing field value with zero value
	stmt, err := s.Db.Prepare("UPDATE students SET name = ?, email = ?, age = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(payload.Name, payload.Email, payload.Age, id)
	if err != nil {
		return err
	}
	return nil

	/* dynamic query to update particular field value

	query := "UPDATE students SET "
	params := []interface{}{}
	if payload.Name != nil {
		query += "name = ?, "
		params = append(params, payload.Name)
	}
	if payload.Email != nil {
		query += "email = ?, "
		params = append(params, payload.Email)
	}
	if payload.Age != nil {
		query += "age = ? "
		params = append(params, payload.Age)
	}
	query = strings.TrimSuffix(query, " ")
	query += " WHERE id = ?"
	params = append(params, id)

	stmt, err := s.Db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(params...)
	if err != nil {
		return err
	}
	return nil
	*/
}
