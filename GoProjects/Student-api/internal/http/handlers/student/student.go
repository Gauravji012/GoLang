package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"reflect"
	"strconv"

	"github.com/Gauravji012/GoLang/GoProjects/Student-api/internal/storage"
	"github.com/Gauravji012/GoLang/GoProjects/Student-api/internal/types"
	"github.com/Gauravji012/GoLang/GoProjects/Student-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

// in go, we dont decode data directly. first we take & serialize that data & then we will use it.
// json data ko struct student(inside types.go) mein serialize krna hai -> for the purpose of use

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("creating a student")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		//handle general error
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// always validate the user request -- dont rely on user  -- verification is use for safe sidee-> good practice
		// request validation -- we can do it manually or use golang package which is playground validator
		// playground packagee link - "go get github.com/go-playground/validator/v10"
		// we apply field validation of struct to use playground package ============

		// request validation
		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		// now we create our student
		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)

		slog.Info("user created successfully", slog.String("userId", fmt.Sprint(lastId)))

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})
		// w.Write([]byte("Welcome to student api"))
	}
}

// after creating new routing end point we handle it using hanlder function

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id") // to access the query parameter using pathvalue go lang inbuilt function
		slog.Info("getting a student", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Info("error while parsing the parameter")
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		student, err := storage.GetStudentById(intId)
		if err != nil {
			slog.Info("error while getting user", slog.String("id", id))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		response.WriteJson(w, http.StatusOK, student)
	}
}

func UpdateStudentById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("getting a student ", slog.String("id", id))
		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Info("error while parsing the parameter")
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

/*
var payloadMap map[string]interface{}
err := json.NewDecoder(r.Body).Decode(&payloadMap)
if err != nil {
    // handle error
}

// Fetch current student from DB by ID (existing data)
existingStudent, err := storage.GetStudentById(id)
if err != nil {
    // handle error
}

// Update only the fields in payloadMap
if name, ok := payloadMap["name"].(string); ok {
    existingStudent.Name = name
}
if email, ok := payloadMap["email"].(string); ok {
    existingStudent.Email = email
}
if age, ok := payloadMap["age"].(float64); ok { // JSON numbers decode to float64
    existingStudent.Age = int(age)
}

// Now update the DB with existingStudent which has updated fields
err = storage.UpdateStudentInDB(id, existingStudent)
if err != nil {
    // handle error
}

*/
		var student types.Student
		// decoding the user request data
		err = json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		err = storage.UpdateStudent(intId, student)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return 
		}
		response.WriteJson(w, http.StatusOK, map[string]string{"message": "student is successfully updated"})
	}
}

func GetList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("getting all students")

		students, err := storage.GetStudents()
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}
		response.WriteJson(w, http.StatusOK, students)
	}
}


func RemoveStudent(storage storage.Storage) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("delete an student", slog.String("id", id), slog.String("type", reflect.TypeOf(id).String()))
		intId, err := strconv.ParseInt(id, 10, 64)
		
		if err != nil {
			slog.Info("error while parsing the parameter")
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		err = storage.DeleteStudentFromDB(intId)
		if err != nil {
			slog.Info("error while delete")
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return 
		}
		response.WriteJson(w, http.StatusOK, map[string]string{"message": "student is successfully deleted from db"})

	}
}