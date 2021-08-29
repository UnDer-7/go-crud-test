package app_error

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AppError struct {
	StatusCode       int    `json:"statusCode,omitempty"`
	UserMessage      string `json:"userMessage,omitempty"`
	DeveloperMessage string `json:"developerMessage,omitempty"`
	OriginalError    error  `json:"originalError,omitempty"`
}

func ThrowInternalServerError(developerMsg string, originErr error) *AppError {
	return &AppError{
		StatusCode:       http.StatusInternalServerError,
		UserMessage:      "Algo deu errado",
		DeveloperMessage: developerMsg,
		OriginalError:    originErr,
	}
}

func ThrowNotFoundError(userMsg, developerMsg string, originErr error) *AppError {
	return &AppError{
		StatusCode:       http.StatusNotFound,
		UserMessage:      userMsg,
		DeveloperMessage: developerMsg,
		OriginalError:    originErr,
	}
}

func (ae AppError) Error() string {
	jsonString, err := json.Marshal(ae)
	if err != nil {
		errMsg := "erro ao serializar json do AppError"
		fmt.Printf(errMsg+"\n%v", err)
		return errMsg
	}

	return string(jsonString)
}
