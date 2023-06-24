package application

import (
	"net/http"
	"strings"

	"github.com/yakob-abada/go-api/go/app/repository"
	"github.com/yakob-abada/go-api/go/app/service"

	"github.com/gin-gonic/gin"
)

type SessionListHandler struct {
	SessionRepository    *repository.SessionRepository
	UserRepository       *repository.UserRepository
	ErrorResponseHandler service.ErrorResponse
	UserAuthorization    *service.UserAuthoriztion
}

func (slh *SessionListHandler) GetActiveList(c *gin.Context) {
	result, err := slh.SessionRepository.FindActive()

	if err != nil {
		c.JSON(slh.ErrorResponseHandler.GenerateResponse(err))
		return
	}

	c.JSON(http.StatusOK, result)
}

func (slh *SessionListHandler) GetList(c *gin.Context) {
	result, err := slh.SessionRepository.FindActive()

	if err != nil {
		c.JSON(slh.ErrorResponseHandler.GenerateResponse(err))
		return
	}

	c.JSON(http.StatusOK, result)
}

func (slh *SessionListHandler) Join(c *gin.Context) {
	claims, err := slh.UserAuthorization.Authorize(c)

	if err != nil {
		c.JSON(slh.ErrorResponseHandler.GenerateResponse(err))
		return
	}

	session, err := slh.SessionRepository.FindById(c.Param("id"))

	// @todo need to be refactored.
	if err != nil {
		c.JSON(slh.ErrorResponseHandler.GenerateResponse(service.NewNotFoundError("session not found")))
		return
	}

	err = slh.SessionRepository.Join(session.Id, claims.UserId)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = service.NewUnprocessableEntityError("user has already joined given session")
		}
		c.JSON(slh.ErrorResponseHandler.GenerateResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user has joined successfully"}) //@todo the way json response handled needs to be refactored.
}
