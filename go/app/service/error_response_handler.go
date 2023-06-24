package service

import (
	"net/http"
	"os"

	"github.com/yakob-abada/go-api/go/app/entity"
)

func NewErrorResponseHandler() *ErrorResponseHandler {
	return &ErrorResponseHandler{}
}

type ErrorResponseHandler struct{}

func (erg *ErrorResponseHandler) GenerateResponse(status uint, e error) *entity.ErrorResponse {
	errorMessage := "Something went wrong on Server, we will fix it"

	if os.Getenv("ENV") == "dev" || http.StatusInternalServerError != status {
		errorMessage = e.Error()
	}

	return &entity.ErrorResponse{
		Error: errorMessage,
	}
}
