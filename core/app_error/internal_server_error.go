package app_error

import (
	"errors"
	"fmt"
	"net/http"
)

type InternalServerError struct {
	StatusCode int
	Err        error
}

func NewInternalServerError(msg string, err error) *InternalServerError {
	StatusCode := http.StatusInternalServerError
	if err != nil {
		return &InternalServerError{
			StatusCode: StatusCode,
			Err:        err,
		}
	}

	return &InternalServerError{
		StatusCode: StatusCode,
		Err:        errors.New(msg),
	}
}

func (r *InternalServerError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}
