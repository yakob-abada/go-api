package application

import (
	"net/http"

	"github.com/yakob-abada/go-api/go/app/repository"
	"github.com/yakob-abada/go-api/go/app/service"

	"github.com/gin-gonic/gin"
)

type ProductListHandler struct {
	Repository           repository.Repository
	ErrorResponseHandler service.ErrorResponse
}

func (plh *ProductListHandler) GetProductList(c *gin.Context) {
	result, err := plh.Repository.FindAll()

	if err != nil {
		c.JSON(plh.ErrorResponseHandler.GenerateResponse(err))
		return
	}

	c.JSON(http.StatusOK, result)
}
