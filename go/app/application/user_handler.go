package application

import (
	"fmt"
	"net/http"

	"github.com/yakob-abada/go-api/go/app/repository"
	"github.com/yakob-abada/go-api/go/app/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Repository           *repository.UserRepository
	ErrorResponseHandler service.ErrorResponse
	UserAuthorization    *service.UserAuthoriztion
}

func (ah *UserHandler) Login(c *gin.Context) {
	var authUser service.AuthUser

	if err := c.ShouldBindJSON(&authUser); err != nil {
		c.JSON(ah.ErrorResponseHandler.GenerateResponse(service.NewBadRequestError("username and/or password are wrong")))
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
