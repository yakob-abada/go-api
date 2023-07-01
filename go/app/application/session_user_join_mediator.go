package application

import (
	"strings"

	"github.com/yakob-abada/go-api/go/app/domain"
	"github.com/yakob-abada/go-api/go/app/entity"
	"github.com/yakob-abada/go-api/go/app/service"
)

type ISessionUserJoinMediator interface {
	Mediate(session *entity.Session, userId int8) error
}

func NewSessionUserJoinMediator(
	sessionRepository domain.SessionRepository,
	activeSessionSpecification *ActiveSessionSpecification,
) *SessionUserJoinMediator {
	return &SessionUserJoinMediator{
		sessionRepository:          sessionRepository,
		activeSessionSpecification: activeSessionSpecification,
	}
}

type SessionUserJoinMediator struct {
	sessionRepository          domain.SessionRepository
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
