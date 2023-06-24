package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

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

type UserAuthoriztion struct {
	jwtKey []byte
}

func NewUserAuthorization(jwtKey []byte, jwtExpirationTime int8) *UserAuthoriztion {
	return &UserAuthoriztion{
		jwtKey: jwtKey,
	}
}

func (ua *UserAuthoriztion) GenerateToken(username string, user_id int8) (*TokenResponse, error) {
	// @todo move constant to config value
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: username,
		UserId:   user_id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(ua.jwtKey)

	if err != nil {
		return nil, fmt.Errorf("failed to generate JWT token")
	}

	return &TokenResponse{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	}, nil
}

func (ua *UserAuthoriztion) Authorize(c *gin.Context) (*Claims, error) {
	bearerToken := c.GetHeader("Authorization")

	if bearerToken == "" {
		return nil, NewUnauthorizedError("user is not authorized")
	}

	token := strings.Replace(bearerToken, "Bearer ", "", 1)

	if token == "" {
		return nil, NewUnauthorizedError("user is not authorized")
	}

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return ua.jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, NewUnauthorizedError("user is not authorized")
		}
		return nil, err
	}

	if !tkn.Valid {
		return nil, NewUnauthorizedError("user is not authorized")
	}

	return claims, nil
}
