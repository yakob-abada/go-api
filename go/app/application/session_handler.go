package application

import (
	"net/http"

	"github.com/yakob-abada/go-api/go/app/domain"
	"github.com/yakob-abada/go-api/go/app/model"
	"github.com/yakob-abada/go-api/go/app/service"

	"github.com/gin-gonic/gin"
)

type SessionHandler struct {
	SessionRepository       domain.SessionRepository
	UserRepository          domain.UserRepository
	ErrorResponseHandler    service.ErrorResponse
	UserAuthorization       domain.UserAuthoriztion
	SessionUserJoinMediator ISessionUserJoinMediator
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
	result, err := slh.SessionRepository.FindAll()

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
