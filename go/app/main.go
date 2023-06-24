package main

import (
	"github.com/yakob-abada/go-api/go/app/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/product/:id", bootstrap.NewProductItemHandler().GetProduct)
	r.GET("/product", bootstrap.NewProductListHandler().GetProductList)
	r.GET("/active-session", bootstrap.NewSessionListHandler().GetActiveList)
	r.GET("/session", bootstrap.NewSessionListHandler().GetList)
	r.POST("/session/:id/join", bootstrap.NewSessionListHandler().Join) //@todo choose put or post
	r.POST("/login", bootstrap.NewAuthHandler().Login)
	r.Run() // listen and serve on 0.0.0.0:8080
}
