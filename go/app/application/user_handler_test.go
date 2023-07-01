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
		var mockAppCrypto = &service.MockAppCrypto{}

		user := entity.User{
			Id:       1,
			Username: "username",
			Password: "hashedPassword",
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
		mockUserRepository.On("FindByUsername", "username").Return(&user, nil)
		mockAppCrypto.On("CompareHashAndPassword", []byte("hashedPassword"), []byte("password")).Return(nil)
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
			Encryption:           mockAppCrypto,
		}

		sut.Login(c)
	})

	t.Run("failed no body request", func(t *testing.T) {
		var mockUserRepository = &repository.MockUserRepository{}
		var mockErrorResponse = &service.MockErrorResponse{}
		var mockValidate = &MockValidate{}
		var mockUserAuthorization = &service.MockUserAuthoriztion{}
		var mockAppCrypto = &service.MockAppCrypto{}

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
			Encryption:           mockAppCrypto,
		}

		sut.Login(c)
	})

	t.Run("failed missing username", func(t *testing.T) {
		var mockUserRepository = &repository.MockUserRepository{}
		var mockErrorResponse = &service.MockErrorResponse{}
		var mockValidate = &MockValidate{}
		var mockUserAuthorization = &service.MockUserAuthoriztion{}
		var mockAppCrypto = &service.MockAppCrypto{}

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
			Encryption:           mockAppCrypto,
		}

		sut.Login(c)
	})

	t.Run("failed username is wrong", func(t *testing.T) {
		var mockUserRepository = &repository.MockUserRepository{}
		var mockErrorResponse = &service.MockErrorResponse{}
		var mockValidate = &MockValidate{}
		var mockUserAuthorization = &service.MockUserAuthoriztion{}
		var mockAppCrypto = &service.MockAppCrypto{}

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		userAuth := service.AuthUser{
			Username: "usernames",
			Password: "password",
		}

		mockValidate.On("Struct", userAuth).Return(nil)
		mockUserRepository.On("FindByUsername", "usernames").Return(nil, fmt.Errorf("username doesn't exist"))
		mockErrorResponse.On("GenerateResponse", fmt.Errorf("username doesn't exist")).Return(500, &model.ErrorResponse{Error: "username doesn't exist"}).Once()

		jsonbytes, _ := json.Marshal(userAuth)
		c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonbytes))

		sut := UserHandler{
			Repository:           mockUserRepository,
			ErrorResponseHandler: mockErrorResponse,
			UserAuthorization:    mockUserAuthorization,
			Validate:             mockValidate,
			Encryption:           mockAppCrypto,
		}

		sut.Login(c)
	})

	t.Run("failed password", func(t *testing.T) {
		var mockUserRepository = &repository.MockUserRepository{}
		var mockErrorResponse = &service.MockErrorResponse{}
		var mockValidate = &MockValidate{}
		var mockUserAuthorization = &service.MockUserAuthoriztion{}
		var mockAppCrypto = &service.MockAppCrypto{}

		user := entity.User{
			Id:       1,
			Username: "username",
			Password: "hashedPassword",
		}
		userAuth := service.AuthUser{
			Username: "username",
			Password: "password",
		}

		mockValidate.On("Struct", userAuth).Return(nil)
		mockUserAuthorization.On("GenerateToken", user.Username, user.Id).Return(nil, fmt.Errorf("failed to generate JWT token"))
		mockUserRepository.On("FindByUsername", "username").Return(&user, nil)
		mockAppCrypto.On("CompareHashAndPassword", []byte("hashedPassword"), []byte("password")).Return(fmt.Errorf("invalid auth"))
		mockErrorResponse.On("GenerateResponse", service.NewUnauthorizedError("Invalid credentials")).Return(405, &model.ErrorResponse{Error: "Invalid credentials"}).Once()
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
			Encryption:           mockAppCrypto,
		}

		sut.Login(c)
	})

	t.Run("failed create token", func(t *testing.T) {
		var mockUserRepository = &repository.MockUserRepository{}
		var mockErrorResponse = &service.MockErrorResponse{}
		var mockValidate = &MockValidate{}
		var mockUserAuthorization = &service.MockUserAuthoriztion{}
		var mockAppCrypto = &service.MockAppCrypto{}

		user := entity.User{
			Id:       1,
			Username: "username",
			Password: "hashedPassword",
		}
		userAuth := service.AuthUser{
			Username: "username",
			Password: "password",
		}

		mockValidate.On("Struct", userAuth).Return(nil)
		mockUserAuthorization.On("GenerateToken", user.Username, user.Id).Return(nil, fmt.Errorf("failed to generate JWT token"))
		mockErrorResponse.On("GenerateResponse", fmt.Errorf("failed to generate JWT token")).Return(500, &model.ErrorResponse{Error: "failed to generate JWT token"}).Once()
		mockUserRepository.On("FindByUsername", "username").Return(&user, nil)
		mockAppCrypto.On("CompareHashAndPassword", []byte("hashedPassword"), []byte("password")).Return(nil)
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
			Encryption:           mockAppCrypto,
		}

		sut.Login(c)
	})
}
