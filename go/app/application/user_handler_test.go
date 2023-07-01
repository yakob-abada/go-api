package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/yakob-abada/go-api/go/app/entity"
	"github.com/yakob-abada/go-api/go/app/model"
	"github.com/yakob-abada/go-api/go/app/repository"
	"github.com/yakob-abada/go-api/go/app/service"
)

func TestLogin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var mockUserRepository = &repository.MockUserRepository{}
		var mockErrorResponse = &service.MockErrorResponse{}
		var mockValidate = &MockValidate{}
		var mockUserAuthorization = &service.MockUserAuthoriztion{}

		user := entity.User{
			Id:       1,
			Username: "username",
		}
		userAuth := service.AuthUser{
			Username: "username",
			Password: "password",
		}

		tokenResponse := service.TokenResponse{
			Name:  "name",
			Value: "test",
		}

		mockValidate.On("Struct", userAuth).Return(nil)
		mockUserAuthorization.On("GenerateToken", user.Username, user.Id).Return(&tokenResponse, nil)
		mockUserRepository.On("FindByUsernameAndPass", "username", "password").Return(&user, nil)
		gin.SetMode(gin.TestMode)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		jsonbytes, _ := json.Marshal(userAuth)

		c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonbytes))

		sut := UserHandler{
			Repository:           mockUserRepository,
			ErrorResponseHandler: mockErrorResponse,
			UserAuthorization:    mockUserAuthorization,
			Validate:             mockValidate,
		}

		sut.Login(c)

	})

	t.Run("failed no body request", func(t *testing.T) {
		var mockUserRepository = &repository.MockUserRepository{}
		var mockErrorResponse = &service.MockErrorResponse{}
		var mockValidate = &MockValidate{}
		var mockUserAuthorization = &service.MockUserAuthoriztion{}

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		mockErrorResponse.On("GenerateResponse", service.NewBadRequestError("JSON body request problem")).Return(400, &model.ErrorResponse{Error: "JSON body request problem"}).Once()

		c.Request, _ = http.NewRequest(http.MethodPost, "/", nil)

		sut := UserHandler{
			Repository:           mockUserRepository,
			ErrorResponseHandler: mockErrorResponse,
			UserAuthorization:    mockUserAuthorization,
			Validate:             mockValidate,
		}

		sut.Login(c)

	})

	t.Run("failed missing username", func(t *testing.T) {
		var mockUserRepository = &repository.MockUserRepository{}
		var mockErrorResponse = &service.MockErrorResponse{}
		var mockValidate = &MockValidate{}
		var mockUserAuthorization = &service.MockUserAuthoriztion{}

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		userAuth := service.AuthUser{
			Password: "password",
		}

		mockValidate.On("Struct", userAuth).Return(fmt.Errorf("Username is missing"))
		mockErrorResponse.On("GenerateResponse", service.NewBadRequestError("Username is missing")).Return(400, &model.ErrorResponse{Error: "Username is missing"}).Once()

		jsonbytes, _ := json.Marshal(userAuth)
		c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonbytes))

		sut := UserHandler{
			Repository:           mockUserRepository,
			ErrorResponseHandler: mockErrorResponse,
			UserAuthorization:    mockUserAuthorization,
			Validate:             mockValidate,
		}

		sut.Login(c)

	})
}
