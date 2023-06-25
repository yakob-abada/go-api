package application

import (
	"time"

	"github.com/yakob-abada/go-api/go/app/entity"
)

func NewActiveSessionSpecification() *ActiveSessionSpecification {
	return &ActiveSessionSpecification{}
}

type ActiveSessionSpecification struct{}

func (ass *ActiveSessionSpecification) IsSatisfied(session *entity.Session) bool {
	return (!session.IsFull) && session.Time.After(time.Now())
}
