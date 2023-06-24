package application

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/yakob-abada/go-api/go/app/repository"
	"github.com/yakob-abada/go-api/go/app/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Repository           *repository.UserRepository
	ErrorResponseHandler service.ErrorResponse
	JwtKey               []byte
}

type Claims struct {
	Username string `json:"username"`
	UserId   int8   `json:"user_id"`
	jwt.RegisteredClaims
}

type TokenResponse struct {
	Name    string    `json:"name"`
	Value   string    `json:"value"`
	Expires time.Time `json:"expires"`
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var authUser AtuhUser

	if err := c.ShouldBindJSON(&authUser); err != nil {
		c.JSON(http.StatusBadRequest, ah.ErrorResponseHandler.GenerateResponse(http.StatusBadRequest, fmt.Errorf("username and/or password are missing")))
		return
	}

	// @todo username and password doesn't exist should be handled
	user, err := ah.Repository.FindByUsernameAndPass(authUser.Username, authUser.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ah.ErrorResponseHandler.GenerateResponse(http.StatusInternalServerError, err))
		return
	}

	// @todo move constant to config value
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: user.Username,
		UserId:   user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(ah.JwtKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ah.ErrorResponseHandler.GenerateResponse(http.StatusInternalServerError, fmt.Errorf("failed to generate JWT token")))
		return
	}

	c.JSON(http.StatusOK, TokenResponse{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
