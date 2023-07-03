package service

import (
	"github.com/stretchr/testify/mock"
	"github.com/yakob-abada/go-api/go/app/domain"
)

type MockErrorResponse struct {
	mock.Mock
}

func (mer *MockErrorResponse) GenerateResponse(err error) (int, *domain.ErrorResponse) {
	args := mer.Called(err)

	return args.Get(0).(int), args.Get(1).(*domain.ErrorResponse)
}

type MockAuthToken struct {
	mock.Mock
}

func (mat *MockAuthToken) GenerateToken(username string, userId int8) (*domain.TokenResponse, error) {
	args := mat.Called(username, userId)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.TokenResponse), args.Error(1)
}

func (mat *MockAuthToken) Validate(token string) (*domain.Claims, error) {
	args := mat.Called(token)

	return args.Get(0).(*domain.Claims), args.Error(1)
}

type MockAppCrypto struct {
	mock.Mock
}

func (mac *MockAppCrypto) GenerateFromPassword(password []byte) ([]byte, error) {
	args := mac.Called(password)

	return args.Get(0).([]byte), args.Error(1)
}

func (mac *MockAppCrypto) CompareHashAndPassword(hashedPassword, password []byte) error {
	args := mac.Called(hashedPassword, password)

	return args.Error(0)
}
