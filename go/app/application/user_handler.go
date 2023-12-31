package application

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yakob-abada/go-api/go/app/domain"
	"github.com/yakob-abada/go-api/go/app/service"
)

type UserHandler struct {
	Repository           domain.IUserRepository
	ErrorResponseHandler service.ErrorResponse
	AuthToken            domain.IAuthToken
	Validate             domain.IValidate
	Encryption           domain.IAppCrypto
}

func (ah *UserHandler) Login(c *gin.Context) {
	var authUser domain.AuthUser

	if err := c.ShouldBindJSON(&authUser); err != nil {
		c.JSON(ah.ErrorResponseHandler.GenerateResponse(service.NewBadRequestError("JSON body request problem")))
		return
	}

	// @todo improve error message using translation
	if err := ah.Validate.Struct(authUser); err != nil {
		c.JSON(ah.ErrorResponseHandler.GenerateResponse(service.NewBadRequestError(err.Error())))
		return
	}

	// @todo username and password doesn't exist should be handled
	user, err := ah.Repository.FindByUsername(authUser.Username)

	if err != nil {
		c.JSON(ah.ErrorResponseHandler.GenerateResponse(err))
		return
	}

	if err = ah.Encryption.CompareHashAndPassword([]byte(user.Password), []byte(authUser.Password)); err != nil {
		c.JSON(ah.ErrorResponseHandler.GenerateResponse(service.NewUnauthorizedError("Invalid credentials")))
		return
	}

	tokenResponse, err := ah.AuthToken.GenerateToken(user.Username, user.Id)

	if err != nil {
		c.JSON(ah.ErrorResponseHandler.GenerateResponse(fmt.Errorf("failed to generate JWT token")))
		return
	}

	c.JSON(http.StatusOK, tokenResponse)
}
