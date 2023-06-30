package application

import (
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yakob-abada/go-api/go/app/entity"
	"github.com/yakob-abada/go-api/go/app/model"
	"github.com/yakob-abada/go-api/go/app/repository"
	"github.com/yakob-abada/go-api/go/app/service"
)

func TestSessionHandlerGetList(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var mockSessionRepository = &repository.MockSessionRepository{}
		var mockUserRepository = &repository.MockUserRepository{}
		var mockErrorResponse = &service.MockErrorResponse{}
		var mockUserAuthorization = &service.MockUserAuthoriztion{}
		var mockSessionUserJoinMediator = &MockSessionUserJoinMediator{}
		var time, _ = time.Parse("2006-01-02 15:04:00", "2023-06-26 07:00:00")

		var sessions = []entity.Session{
			{
				Id:       1,
				Time:     time,
				Name:     "Session name",
				Duration: 30,
				IsFull:   false,
			},
		}

		mockSessionRepository.On("FindAll").Return(sessions, nil)
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		sut := &SessionHandler{
			SessionRepository:       mockSessionRepository,
			UserRepository:          mockUserRepository,
			ErrorResponseHandler:    mockErrorResponse,
			UserAuthorization:       mockUserAuthorization,
			SessionUserJoinMediator: mockSessionUserJoinMediator,
		}

		sut.GetList(c)
		mockSessionRepository.AssertNumberOfCalls(t, "FindAll", 1)
	})

	t.Run("fail", func(t *testing.T) {
		var mockSessionRepository = &repository.MockSessionRepository{}
		var mockUserRepository = &repository.MockUserRepository{}
		var mockErrorResponse = &service.MockErrorResponse{}
		var mockUserAuthorization = &service.MockUserAuthoriztion{}
		var mockSessionUserJoinMediator = &MockSessionUserJoinMediator{}

		var sessions []entity.Session
		mockSessionRepository.On("FindAll").Return(sessions, fmt.Errorf("something went wrongss"))
		mockErrorResponse.On("GenerateResponse", fmt.Errorf("something went wrongss")).Return(500, &model.ErrorResponse{Error: "something went wrong"}).Once()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		sut := &SessionHandler{
			SessionRepository:       mockSessionRepository,
			UserRepository:          mockUserRepository,
			ErrorResponseHandler:    mockErrorResponse,
			UserAuthorization:       mockUserAuthorization,
			SessionUserJoinMediator: mockSessionUserJoinMediator,
		}

		sut.GetList(c)

		mockSessionRepository.AssertNumberOfCalls(t, "FindAll", 1)
		mockErrorResponse.AssertNumberOfCalls(t, "GenerateResponse", 1)
	})
}
