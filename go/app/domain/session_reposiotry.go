package domain

import "github.com/yakob-abada/go-api/go/app/entity"

type ISessionRepository interface {
	FindById(id string) (*entity.Session, error)
	FindAll() ([]entity.Session, error)
	FindActive() ([]entity.Session, error)
	Join(sessionId int8, userId int8) error
	SetSessionIsFullSatistfaction(sessionId int8) error
}
