package service

import "github.com/yakob-abada/go-api/go/app/domain"

type ErrorResponse interface {
	GenerateResponse(error) (int, *domain.ErrorResponse)
}
