package bootstrap

import (
	"fmt"
	"os"

	"github.com/yakob-abada/go-api/go/app/application"
	"github.com/yakob-abada/go-api/go/app/repository"
)

func NewProductListHandler() *application.ProductListHandler {
	return &application.ProductListHandler{
		Repository: repository.NewProductRepository(
			repository.NewMysqlConnection(
				os.Getenv("DATABASE_USERNAME"),
				os.Getenv("DATABASE_PASSWORD"),
				fmt.Sprintf("%s:%s", os.Getenv("DATABASE_HOST"), "3306"),
				os.Getenv("DATABASE_NAME"),
			),
		),
	}
}
