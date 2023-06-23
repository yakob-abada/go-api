package main

import (
	"github.com/yakob-abada/go-api/go/app/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/product/:id", bootstrap.NewProductItemHandler().GetProduct)
	r.GET("/product", bootstrap.NewProductListHandler().GetProductList)
	r.Run() // listen and serve on 0.0.0.0:8080
}
