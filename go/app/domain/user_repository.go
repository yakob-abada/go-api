package domain

import "github.com/yakob-abada/go-api/go/app/entity"

type IUserRepository interface {
	FindByUsername(username string) (*entity.User, error)
}
