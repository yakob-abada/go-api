package application

import (
	"github.com/yakob-abada/go-api/go/app/entity"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mr *MockRepository) FindById(id string) (*entity.Product, error) {
	args := mr.Called(id)

	return args.Get(0).(*entity.Product), args.Error(1)
}
func (mr *MockRepository) FindAll() (*[]entity.Product, error) {
	args := mr.Called()

	return args.Get(0).(*[]entity.Product), args.Error(1)
}

type MockErrorResponse struct {
	mock.Mock
}

func (mer *MockErrorResponse) GenerateResponse(error) *entity.ErrorResponse {
	args := mer.Called()

	return args.Get(0).(*entity.ErrorResponse)
}
