package server

import "net/http"

type ServerError struct {
	Message string `json:"message"`
	Code    int    `json:"-"`
}

func NewError(message string, code int) *ServerError {
	return &ServerError{Message: message, Code: code}
}

func (e *ServerError) Json(rw http.ResponseWriter) {
	Json(rw, e, e.Code)
}

func (e *ServerError) Error() string {
	return e.Message
}
