package service

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/yakob-abada/go-api/go/app/model"
)

type MockErrorResponse struct {
	mock.Mock
}

func (mer *MockErrorResponse) GenerateResponse(err error) (int, *model.ErrorResponse) {
	args := mer.Called(err)

	return args.Get(0).(int), args.Get(1).(*model.ErrorResponse)
}

type MockUserAuthoriztion struct {
	mock.Mock
}

func (mua *MockUserAuthoriztion) GenerateToken(username string, userId int8) (*TokenResponse, error) {
	args := mua.Called(username, userId)

	return args.Get(0).(*TokenResponse), args.Error(1)
}

func (mua *MockUserAuthoriztion) Authorize(c *gin.Context) (*Claims, error) {
	args := mua.Called(c)

	return args.Get(0).(*Claims), args.Error(1)
}
