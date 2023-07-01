package repository

import (
	"github.com/yakob-abada/go-api/go/app/entity"

	"github.com/stretchr/testify/mock"
)

type MockSessionRepository struct {
	mock.Mock
}

func (msr *MockSessionRepository) FindById(id string) (*entity.Session, error) {
	args := msr.Called(id)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.Session), args.Error(1)
}

func (msr *MockSessionRepository) FindAll() ([]entity.Session, error) {
	args := msr.Called()

	return args.Get(0).([]entity.Session), args.Error(1)
}

func (msr *MockSessionRepository) FindActive() ([]entity.Session, error) {
	args := msr.Called()

	return args.Get(0).([]entity.Session), args.Error(1)
}

func (msr *MockSessionRepository) Join(sessionId int8, userId int8) error {
	args := msr.Called(sessionId, userId)

	return args.Error(1)
}

func (msr *MockSessionRepository) SetSessionIsFullSatistfaction(sessionId int8) error {
	args := msr.Called(sessionId)

	return args.Error(1)
}

type MockUserRepository struct {
	mock.Mock
}

func (mur *MockUserRepository) FindByUsername(username string) (*entity.User, error) {
	args := mur.Called(username)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.User), args.Error(1)
}
