package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/geekyharsh05/students-api/internal/storage"
	"github.com/geekyharsh05/students-api/internal/types"
	"github.com/geekyharsh05/students-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Creating a New Student")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("body is empty")))
			return
		}

		if err != nil {
			slog.Error("Failed to decode body", slog.String("error", err.Error()))
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// request validation
		if err := validator.New().Struct(student); err != nil {
			validateErrors := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrors))
			return
		}

		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)

		if err != nil {
			slog.Error("Failed to create student", slog.String("error", err.Error()))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		slog.Info("Getting a Student", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Error("Failed to parse id", slog.String("error", err.Error()))
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		student, err := storage.GetStudentById(intId)
		if err != nil {
			slog.Error("Failed to get student", slog.String("error", err.Error()))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, student)
	}
}

func GetAllStudents(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Getting all students")

		students, err := storage.GetStudents()
		if err != nil {
			slog.Error("Failed to get all students", slog.String("error", err.Error()))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, students)
	}
}

func UpdateAStudent(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Updating a Student")

		// Get the ID from the URL
		id := r.PathValue("id")
		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Error("Failed to parse id", slog.String("error", err.Error()))
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		var student types.Student

		err = json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("body is empty")))
			return
		}

		if err != nil {
			slog.Error("Failed to decode body", slog.String("error", err.Error()))
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// request validation
		if err := validator.New().Struct(student); err != nil {
			validateErrors := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrors))
			return
		}

		// Check if the student data is already up to date
		existingStudent, err := storage.GetStudentById(intId)
		if err != nil {
			slog.Error("Failed to get existing student", slog.String("error", err.Error()))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		if existingStudent.Name == student.Name && existingStudent.Email == student.Email && existingStudent.Age == student.Age {
			response.WriteJson(w, http.StatusConflict, response.GeneralError(fmt.Errorf("no changes detected, update not allowed")))
			return
		}

		err = storage.UpdateStudent(
			intId,
			student.Name,
			student.Email,
			student.Age,
		)

		if err != nil {
			slog.Error("Failed to update student", slog.String("error", err.Error()))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, map[string]string{"status": "student updated successfully"})
	}
}

func DeleteById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		slog.Info("Deleting a Student", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Error("Failed to parse id", slog.String("error", err.Error()))
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		err = storage.DeleteStudentById(intId)
		if err != nil {
			slog.Error("Failed to delete student", slog.String("error", err.Error()))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, map[string]string{"status": "student deleted successfully"})
	}
}