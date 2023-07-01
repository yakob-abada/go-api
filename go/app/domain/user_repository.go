package domain

import "github.com/yakob-abada/go-api/go/app/entity"

type UserRepository interface {
	FindByUsername(username string) (*entity.User, error)
}
