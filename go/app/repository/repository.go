package repository

import "github.com/yakob-abada/go-api/go/app/entity"

type Repository interface {
	FindById(sku string) (*entity.Product, error)
	FindAll() (*[]entity.Product, error)
}
