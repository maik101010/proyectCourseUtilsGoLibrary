package rest_errors

import (
	"errors"
	"net/http"
)

//RestError struct
type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"code"`
	Error   string `json:"error"`
	Causes []interface{} `json:"causes"`

}

func NewError(msg string) error {
	return errors.New(msg)
}

//NewBadRequestError error function
func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

//NewNotFoundError error function
func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

//NewInternalServerError error function
func NewInternalServerError(message string, err error) *RestError {
	result:= &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
	if err !=nil{
		result.Causes = append(result.Causes,err.Error())
	}
	return result
}

