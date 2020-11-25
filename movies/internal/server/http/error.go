package http

import (
	"net/http"
)

type RestError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewInternalServerError(err error) RestError {
	return RestError{
		Code:    http.StatusInternalServerError,
		Message: "It is not possible to serve you what you are asking for at the moment. It is our fault but life's a bitch.",
	}
}

func (re RestError) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
