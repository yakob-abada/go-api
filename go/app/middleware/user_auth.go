package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yakob-abada/go-api/go/app/domain"
	"github.com/yakob-abada/go-api/go/app/service"
)

type authHeader struct {
	IDToken string `header:"Authorization"`
}

type UserAuth struct {
	AuthToken            domain.IAuthToken
	ErrorResponseHandler service.ErrorResponse
}

func (ua *UserAuth) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}

		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(ua.ErrorResponseHandler.GenerateResponse(service.NewBadRequestError("Body request issue")))
			c.Abort()
			return
		}

		idTokenHeader := strings.Split(h.IDToken, "Bearer ")

		if len(idTokenHeader) < 2 {
			c.JSON(ua.ErrorResponseHandler.GenerateResponse(service.NewUnauthorizedError("user is not authorized")))
			c.Abort()
			return
		}

		cliams, err := ua.AuthToken.Validate(idTokenHeader[1])

		if err != nil {
			c.JSON(ua.ErrorResponseHandler.GenerateResponse(service.NewUnauthorizedError("provided token is not valid")))
			c.Abort()
			return
		}

		c.Set("cliams", *cliams)

		c.Next()
	}
}
