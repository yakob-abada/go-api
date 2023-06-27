package application

import (
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/yakob-abada/go-api/go/app/entity"
	"github.com/yakob-abada/go-api/go/app/repository"
	"github.com/yakob-abada/go-api/go/app/service"

	"github.com/gin-gonic/gin"
)

func TestSessionHandler(t *testing.T) {
	time, _ := time.Parse("2006-01-02 15:04:00", "2023-06-26 07:00:00")

	sessions := []entity.Session{
		{
			Id:       1,
			Time:     time,
			Name:     "Session name",
			Duration: 30,
			IsFull:   false,
		},
	}

	mockSessionRepository := &repository.MockSessionRepository{}
	mockUserRepository := &repository.MockUserRepository{}
	mockErrorResponse := &service.MockErrorResponse{}
	mockUserAuthorization := &service.MockUserAuthoriztion{}
	mockSessionUserJoinMediator := &MockSessionUserJoinMediator{}

	t.Run("GetList", func(t *testing.T) {
		mockSessionRepository.On("FindAll").Return(&sessions, nil)
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
	})

	t.Run("GetListFail", func(t *testing.T) {
		mockSessionRepository.On("FindAll").Return(nil, fmt.Errorf("something went wrong"))
		mockErrorResponse.On("GenerateResponse", 500, fmt.Errorf("something went wrong"))

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
	})
}
