package domain

import "github.com/gin-gonic/gin"

type UserAuthoriztion interface {
	GenerateToken(username string, userId int8) (*TokenResponse, error)
	Authorize(c *gin.Context) (*Claims, error)
}
