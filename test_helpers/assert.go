package test_helpers

import (
	"my-tracking-list-backend/core/app_error"
	"testing"
)

type Asserts struct {
	t testing.TB
}

func NewAssert(t testing.TB) *Asserts {
	return &Asserts{t}
}

func (a Asserts) NotExpectedError(err error) {
	t := a.t
	t.Helper()

	if err != nil {
		t.Fatalf("Got an error but didn't want one. Err: %s", err)
	}
}

func (a Asserts) ExpectedError(err error) {
	t := a.t
	t.Helper()

	if err == nil {
		t.Fatal("Expected an error to occur, but it didn't")
	}
}

func (a Asserts) ExpectedErrorStatusCode(expectedStatusCode int, err *app_error.AppError) {
	t := a.t
	t.Helper()

	if err.StatusCode != expectedStatusCode {
		t.Fatalf("Expected the error status code to be %d but it is %d",expectedStatusCode, err.StatusCode)
	}
}