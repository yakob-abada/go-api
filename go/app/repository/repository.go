package repository

import "github.com/yakob-abada/go-api/entity"

type Repository interface {
	FindById(sku string) (*entity.Product, error)
	FindAll() (*[]entity.Product, error)
}
