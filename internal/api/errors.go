package api

import (
	"errors"
	"net/http"
)

// HTTPError is an augmented error with a HTTP status code.
type HTTPError struct {
	StatusCode int
	error
}

// Error implements the error interface.
func (e *HTTPError) Error() string {
	return e.error.Error()
}

func NewMethodNotAllowed(method string) *HTTPError {
	return &HTTPError{http.StatusMethodNotAllowed, errors.New(`Method is not allowed:"` + method + `"`)}
}
