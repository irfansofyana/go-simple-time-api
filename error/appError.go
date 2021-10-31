package error

import "net/http"

type AppError struct {
	Message string `json:"error_message"`
	Code    int    `json:",omitempty"`
}

func NewInvalidTimezoneError() *AppError {
	return &AppError{
		Message: "invalid timezone",
		Code:    http.StatusNotFound,
	}
}

func NewUndefinedEndpointError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func (err *AppError) AsMessage() *AppError {
	return &AppError{
		Message: err.Message,
	}
}
