package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/geekyharsh05/students-api/internal/types"
	"github.com/go-playground/validator/v10"
)

const (
	StatusOk    = "OK"
	StatusError = "ERROR"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) types.Response {
	return types.Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) types.Response {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
			case "required":
				errMsgs = append(errMsgs, fmt.Sprintf("%s is required", err.Field()))
			default:
				errMsgs = append(errMsgs, fmt.Sprintf("%s is invalid", err.Field()))
		}
	}

	return types.Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}
}