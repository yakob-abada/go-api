package bootstrap

import (
	"fmt"
	"os"

	"github.com/yakob-abada/go-api/go/app/application"
	"github.com/yakob-abada/go-api/go/app/repository"
	"github.com/yakob-abada/go-api/go/app/service"
)

func NewAuthHandler() *application.AuthHandler {
	return &application.AuthHandler{
		Repository: repository.NewUserRepository(
			repository.NewMysqlConnection(
				os.Getenv("DATABASE_USERNAME"),
				os.Getenv("DATABASE_PASSWORD"),
				fmt.Sprintf("%s:%s", os.Getenv("DATABASE_HOST"), "3306"),
				os.Getenv("DATABASE_NAME"),
			),
		),
		ErrorResponseHandler: service.NewErrorResponseHandler(),
		JwtKey:               []byte("jwt_secret_key"),
	}
}
