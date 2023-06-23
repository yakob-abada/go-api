package application

import (
	"net/http"

	"github.com/yakob-abada/go-api/go/app/repository"

	"github.com/gin-gonic/gin"
)

type ProductListHandler struct {
	Repository repository.Repository
}

func (plh *ProductListHandler) GetProductList(c *gin.Context) {
	result, _ := plh.Repository.FindAll()

	c.JSON(http.StatusOK, result)
}
