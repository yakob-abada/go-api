package service

import "github.com/yakob-abada/go-api/go/app/entity"

type ErrorResponse interface {
	GenerateResponse(uint, error) *entity.ErrorResponse
}
