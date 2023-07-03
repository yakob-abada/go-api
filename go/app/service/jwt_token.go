package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/yakob-abada/go-api/go/app/domain"
)

type JwtToken struct {
	jwtKey            []byte
	jwtExpirationTime int8
}

func NewJwtToken(jwtKey []byte, jwtExpirationTime int8) *JwtToken {
	return &JwtToken{
		jwtKey:            jwtKey,
		jwtExpirationTime: jwtExpirationTime,
	}
}

func (jt *JwtToken) GenerateToken(username string, userId int8) (*domain.TokenResponse, error) {
	// @todo move constant to config value
	expirationTime := time.Now().Add(time.Duration(jt.jwtExpirationTime) * time.Minute)

	claims := &domain.Claims{
		Username: username,
		UserId:   userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jt.jwtKey)

	if err != nil {
		return nil, fmt.Errorf("failed to generate JWT token")
	}

	return &domain.TokenResponse{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	}, nil
}

func (jt *JwtToken) Validate(token string) (*domain.Claims, error) {
	claims := &domain.Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jt.jwtKey, nil
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
