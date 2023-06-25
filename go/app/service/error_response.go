package service

import (
	"github.com/yakob-abada/go-api/go/app/model"
)

type ErrorResponse interface {
	GenerateResponse(error) (int, *model.ErrorResponse)
}
