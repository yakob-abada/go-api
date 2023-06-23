package application

import (
	"net/http"

	"github.com/yakob-abada/go-api/repository"

	"github.com/gin-gonic/gin"
)

type ProductItemHandler struct {
	Repository repository.Repository
}

func (pth *ProductItemHandler) GetProduct(c *gin.Context) {
	result, _ := pth.Repository.FindById(c.Param("id"))

	c.JSON(http.StatusOK, result)
}
