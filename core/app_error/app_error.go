package app_error

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AppError struct {
	StatusCode       int    `json:"statusCode"`
	UserMessage      string `json:"userMessage"`
	DeveloperMessage string `json:"developerMessage"`
	OriginalError    error  `json:"-"`
}

func ThrowInternalServerError(developerMsg string, originErr error) *AppError {
	return &AppError{
		StatusCode:       http.StatusInternalServerError,
		UserMessage:      "Algo deu errado",
		DeveloperMessage: developerMsg,
		OriginalError:    originErr,
	}
}

func ThrowBadRequestError(usrMsg, developerMsg string, originErr error) *AppError {
	return &AppError{
		StatusCode:       http.StatusBadRequest,
		UserMessage:      usrMsg,
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

func ThrowBusinessError(usrMsg, devMsg string) *AppError {
	return &AppError{
		StatusCode:       http.StatusUnprocessableEntity,
		UserMessage:      usrMsg,
		DeveloperMessage: devMsg,
		OriginalError:    nil,
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
