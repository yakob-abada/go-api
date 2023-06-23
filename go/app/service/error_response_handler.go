package service

import (
	"os"

	"github.com/yakob-abada/go-api/go/app/entity"
)

func NewErrorResponseHandler() *ErrorResponseHandler {
	return &ErrorResponseHandler{}
}

type ErrorResponseHandler struct{}

func (erg *ErrorResponseHandler) GenerateResponse(e error) *entity.ErrorResponse {
	errorMessage := "Something went wrong on Server, we will fix it"

	if os.Getenv("ENV") == "dev" {
		errorMessage = e.Error()
	}

	return &entity.ErrorResponse{
		Error: errorMessage,
	}
}
