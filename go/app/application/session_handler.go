package application

import (
	"net/http"

	"github.com/yakob-abada/go-api/go/app/model"
	"github.com/yakob-abada/go-api/go/app/repository"
	"github.com/yakob-abada/go-api/go/app/service"

	"github.com/gin-gonic/gin"
)

type SessionHandler struct {
	SessionRepository       *repository.SessionRepository
	UserRepository          *repository.UserRepository
	ErrorResponseHandler    service.ErrorResponse
	UserAuthorization       *service.UserAuthoriztion
	SessionUserJoinMediator *SessionUserJoinMediator
}

func (slh *SessionHandler) GetActiveList(c *gin.Context) {
	result, err := slh.SessionRepository.FindActive()

	if err != nil {
		c.JSON(slh.ErrorResponseHandler.GenerateResponse(err))
		return
	}

	c.JSON(http.StatusOK, result)
}

func (slh *SessionHandler) GetList(c *gin.Context) {
	result, err := slh.SessionRepository.FindActive()

	if err != nil {
		c.JSON(slh.ErrorResponseHandler.GenerateResponse(err))
		return
	}

	c.JSON(http.StatusOK, result)
}

func (slh *SessionHandler) Join(c *gin.Context) {
	claims, err := slh.UserAuthorization.Authorize(c)

	if err != nil {
		c.JSON(slh.ErrorResponseHandler.GenerateResponse(err))
		return
	}

	session, err := slh.SessionRepository.FindById(c.Param("id"))

	// @todo need to be refactored.
	if err != nil {
		c.JSON(slh.ErrorResponseHandler.GenerateResponse(err))
		return
	}

	err = slh.SessionUserJoinMediator.Mediate(session, claims.UserId)

	if err != nil {
		c.JSON(slh.ErrorResponseHandler.GenerateResponse(err))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "user has joined successfully"})
}
