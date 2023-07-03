package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yakob-abada/go-api/go/app/domain"
	"github.com/yakob-abada/go-api/go/app/service"
)

func TestAuthUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	validTokenHeader := "validTokenString"
	invalidTokenHeader := "invalidTokenString"
	t.Run("Success", func(t *testing.T) {
		mockAuthToken := &service.MockAuthToken{}
		mockErrorResponse := &service.MockErrorResponse{}

		claims := &domain.Claims{
			Username: "Jakob",
			UserId:   1,
		}

		mockAuthToken.On("Validate", validTokenHeader).Return(claims, nil)

		sut := UserAuth{
			AuthToken:            mockAuthToken,
			ErrorResponseHandler: mockErrorResponse,
		}

		var contextClaims *domain.Claims
		rr := httptest.NewRecorder()

		_, r := gin.CreateTestContext(rr)

		r.GET("/me", sut.Authenticate(), func(c *gin.Context) {
			contextKeyVal, _ := c.Get("cliams")
			contextClaims = contextKeyVal.(*domain.Claims)
		})

		request, _ := http.NewRequest(http.MethodGet, "/me", http.NoBody)

		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", validTokenHeader))
		r.ServeHTTP(rr, request)

		assert.Equal(t, contextClaims, claims)
	})

	t.Run("failed missing headers", func(t *testing.T) {
		mockAuthToken := &service.MockAuthToken{}
		mockErrorResponse := &service.MockErrorResponse{}

		claims := &domain.Claims{
			Username: "Jakob",
			UserId:   1,
		}

		mockAuthToken.On("Validate", validTokenHeader).Return(claims, nil)
		err := service.NewUnauthorizedError("user is not authorized")
		mockErrorResponse.On("GenerateResponse", err).Return(402, &domain.ErrorResponse{Error: "session is not available to join"}).Once()

		sut := UserAuth{
			AuthToken:            mockAuthToken,
			ErrorResponseHandler: mockErrorResponse,
		}

		rr := httptest.NewRecorder()

		_, r := gin.CreateTestContext(rr)

		r.GET("/me", sut.Authenticate(), func(c *gin.Context) {
		})

		request, _ := http.NewRequest(http.MethodGet, "/me", http.NoBody)

		r.ServeHTTP(rr, request)
	})

	t.Run("failed token validation", func(t *testing.T) {
		mockAuthToken := &service.MockAuthToken{}
		mockErrorResponse := &service.MockErrorResponse{}

		err := service.NewUnauthorizedError("provided token is not valid")
		var claims *domain.Claims
		mockAuthToken.On("Validate", invalidTokenHeader).Return(claims, err)
		mockErrorResponse.On("GenerateResponse", err).Return(402, &domain.ErrorResponse{Error: "provided token is not valid"}).Once()

		sut := UserAuth{
			AuthToken:            mockAuthToken,
			ErrorResponseHandler: mockErrorResponse,
		}

		rr := httptest.NewRecorder()

		_, r := gin.CreateTestContext(rr)

		r.GET("/me", sut.Authenticate(), func(c *gin.Context) {
		})

		request, _ := http.NewRequest(http.MethodGet, "/me", http.NoBody)

		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", invalidTokenHeader))
		r.ServeHTTP(rr, request)
	})
}
