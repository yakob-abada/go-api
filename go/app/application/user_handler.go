package application

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yakob-abada/go-api/go/app/domain"
	"github.com/yakob-abada/go-api/go/app/repository"
	"github.com/yakob-abada/go-api/go/app/service"
)

type UserHandler struct {
	Repository           repository.IUserRepository
	ErrorResponseHandler service.ErrorResponse
	UserAuthorization    service.IUserAuthoriztion
	Validate             domain.Validate
}

func (ah *UserHandler) Login(c *gin.Context) {
	var authUser service.AuthUser

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
	user, err := ah.Repository.FindByUsernameAndPass(authUser.Username, authUser.Password)

	if err != nil {
		c.JSON(ah.ErrorResponseHandler.GenerateResponse(err))
		return
	}

	tokenResponse, err := ah.UserAuthorization.GenerateToken(user.Username, user.Id)

	if err != nil {
		c.JSON(ah.ErrorResponseHandler.GenerateResponse(fmt.Errorf("failed to generate JWT token")))
		return
	}

	c.JSON(http.StatusOK, tokenResponse)
}
