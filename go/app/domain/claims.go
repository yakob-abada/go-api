package domain

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Username string `json:"username"`
	UserId   int8   `json:"user_id"`
	jwt.RegisteredClaims
}
