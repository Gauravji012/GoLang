package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/Gauravji012/GoLang/GoProjects/Student-api/internal/types"
	"github.com/Gauravji012/GoLang/GoProjects/Student-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

// in go, we dont decode data directly. first we take & serialize that data & then we will use it.
// json data ko struct student(inside types.go) mein serialize krna hai -> for the purpose of use

func New() http.HandlerFunc {
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

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
		// w.Write([]byte("Welcome to student api"))
	}
}
