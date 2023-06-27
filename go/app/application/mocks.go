package application

import (
	"github.com/stretchr/testify/mock"
	"github.com/yakob-abada/go-api/go/app/entity"
)

type MockSessionUserJoinMediator struct {
	mock.Mock
}

func (msujm *MockSessionUserJoinMediator) Mediate(session *entity.Session, userId int8) error {
	args := msujm.Called(session, userId)

	return args.Error(0)
}
