package service

import (
	"errors"
	"net/http"
	"os"

	"github.com/yakob-abada/go-api/go/app/model"
)

func NewErrorResponseHandler() *ErrorResponseHandler {
	return &ErrorResponseHandler{}
}

type ErrorResponseHandler struct{}

func (erg *ErrorResponseHandler) GenerateResponse(err error) (int, *model.ErrorResponse) {
	errorMessage := err.Error()
	var statusCode int = http.StatusInternalServerError

	var internalServerError *InternalServerError
	if os.Getenv("ENV") == "prod" && errors.As(err, &internalServerError) {
		errorMessage = "Something went wrong on Server, we will fix it"
	}

	var unauthorizedError *UnauthorizedError
	if errors.As(err, &unauthorizedError) {
		statusCode = http.StatusUnauthorized
	}

	var badRequestError *BadRequestError
	if errors.As(err, &badRequestError) {
		statusCode = http.StatusBadRequest
	}

	var unprocessableEntityError *UnprocessableEntityError
	if errors.As(err, &unprocessableEntityError) {
		statusCode = http.StatusUnprocessableEntity
	}

	return statusCode, &model.ErrorResponse{
		Error: errorMessage,
	}
}
