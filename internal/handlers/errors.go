package handlers

import (
	"net/http"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

const (
	ErrInvalidJSON       = "invalid_json"
	ErrInvalidParameters = "invalid_parameters"
	ErrNotFound          = "not_found"
)

func WriteError(w http.ResponseWriter, status int, code string, err error) {
	WriteJSON(w, status, ErrorResponse{
		Error:   code,
		Message: err.Error(),
	})
}
