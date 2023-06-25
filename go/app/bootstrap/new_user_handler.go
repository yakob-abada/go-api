package bootstrap

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/yakob-abada/go-api/go/app/application"
	"github.com/yakob-abada/go-api/go/app/repository"
	"github.com/yakob-abada/go-api/go/app/service"
)

func NewUserHandler() *application.UserHandler {
	return &application.UserHandler{
		Repository: repository.NewUserRepository(
			repository.NewMysqlConnection(
				os.Getenv("DATABASE_USERNAME"),
				os.Getenv("DATABASE_PASSWORD"),
				fmt.Sprintf("%s:%s", os.Getenv("DATABASE_HOST"), "3306"),
				os.Getenv("DATABASE_NAME"),
			),
		),
		ErrorResponseHandler: service.NewErrorResponseHandler(),
		UserAuthorization: service.NewUserAuthorization(
			[]byte(os.Getenv("JWT_SECRET_KEY")), 8,
		),
		Validate: validator.New(),
	}
}