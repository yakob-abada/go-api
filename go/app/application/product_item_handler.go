package application

import (
	"net/http"

	"github.com/yakob-abada/go-api/go/app/repository"

	"github.com/gin-gonic/gin"
)

type ProductItemHandler struct {
	Repository repository.Repository
}

func (pth *ProductItemHandler) GetProduct(c *gin.Context) {
	result, err := pth.Repository.FindById(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, result)
}
