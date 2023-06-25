package application

import (
	"strings"

	"github.com/yakob-abada/go-api/go/app/entity"
	"github.com/yakob-abada/go-api/go/app/repository"
	"github.com/yakob-abada/go-api/go/app/service"
)

func NewSessionUserJoinMediator(
	sessionRepository *repository.SessionRepository,
	activeSessionSpecification *ActiveSessionSpecification,
) *SessionUserJoinMediator {
	return &SessionUserJoinMediator{
		sessionRepository:          sessionRepository,
		activeSessionSpecification: activeSessionSpecification,
	}
}

type SessionUserJoinMediator struct {
	sessionRepository          *repository.SessionRepository
	activeSessionSpecification *ActiveSessionSpecification
}

func (suj *SessionUserJoinMediator) Mediate(session *entity.Session, userId int8) error {
	if !suj.activeSessionSpecification.IsSatisfied(session) {
		return service.NewUnprocessableEntityError("session is not available to join")
	}

	// @todo make it transaction process.
	err := suj.sessionRepository.Join(session.Id, userId)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = service.NewUnprocessableEntityError("user has already joined given session")
		}

		return err
	}

	err = suj.sessionRepository.SetSessionIsFullSatistfaction(session.Id)

	return err
}
