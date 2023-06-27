package main

import (
	"github.com/yakob-abada/go-api/go/app/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/active-session", bootstrap.NewSessionHandler().GetActiveList)
	r.GET("/session", bootstrap.NewSessionHandler().GetList)
	r.POST("/session/:id/join", bootstrap.NewSessionHandler().Join) //@todo choose put or post
	r.POST("/login", bootstrap.NewUserHandler().Login)
	r.Run() // listen and serve on 0.0.0.0:8080
}
