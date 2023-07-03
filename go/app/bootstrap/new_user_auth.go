package bootstrap

import (
	"os"

	"github.com/yakob-abada/go-api/go/app/middleware"
	"github.com/yakob-abada/go-api/go/app/service"
)

func NewUserAuth() *middleware.UserAuth {
	return &middleware.UserAuth{
		AuthToken: service.NewJwtToken(
			[]byte(os.Getenv("JWT_SECRET_KEY")), 8,
		),
		ErrorResponseHandler: service.NewErrorResponseHandler(),
	}
}
