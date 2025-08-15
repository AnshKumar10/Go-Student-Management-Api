package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/AnshKumar10/Go-Student-Management-Api/internal/storage"
	"github.com/AnshKumar10/Go-Student-Management-Api/internal/types"
	"github.com/AnshKumar10/Go-Student-Management-Api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		slog.Info("Creating new student")

		lastId, err := storage.CreateStudent(student.Name, student.Email, student.Age)

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}

		slog.Info("Student created successfully")

		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})

	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		slog.Info("Getting student with id", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		}

		student, err := storage.GetStudentById(intId)

		if err != nil {
			slog.Info("Error getting student with id", slog.String("id", id))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusCreated, student)

	}
}

func UpdateById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		id := r.PathValue("id")

		slog.Info("Updating student with id", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("invalid id format")))
			return
		}

		err = json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		rowsAffected, err := storage.UpdateStudentById(intId, student.Name, student.Email, student.Age)

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		slog.Info("Student updated successfully")

		response.WriteJson(w, http.StatusOK, map[string]int64{"updated": rowsAffected})
	}
}

func DeleteById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("Deleting student with id", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("invalid id format")))
			return
		}

		success, err := storage.DeleteStudentById(intId)

		if err != nil {
			slog.Error("Error deleting student", slog.String("id", id), slog.String("error", err.Error()))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		if !success {
			response.WriteJson(w, http.StatusNotFound, response.GeneralError(fmt.Errorf("student not found")))
			return
		}

		slog.Info("Student deleted successfully", slog.String("id", id))

		response.WriteJson(w, http.StatusOK, map[string]string{
			"message": fmt.Sprintf("student with id %s deleted", id),
		})
	}
}

func GetAll(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("Getting all students")

		students, err := storage.GetAllStudents()

		if err != nil {
			slog.Info("Error getting students with id")
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusCreated, students)

	}
}
