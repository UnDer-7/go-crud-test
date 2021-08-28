package app_error

import (
	"errors"
	"fmt"
	"net/http"
)

type NotFoundError struct {
	StatusCode int
	Err        error
}

func (err NotFoundError) Error() string {
	return fmt.Sprintf("status %d: err %v", err.StatusCode, err.Err)
}

func NewNotFoundError(msg string, err error) *NotFoundError {
	statusCode := http.StatusNotFound
	if err != nil {
		return &NotFoundError{
			StatusCode: statusCode,
			Err:        err,
		}
	}

	return &NotFoundError{
		StatusCode: statusCode,
		Err:        errors.New(msg),
	}
}
