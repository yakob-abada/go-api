package application

import (
	"net/http/httptest"
	"testing"

	"github.com/yakob-abada/go-api/go/app/entity"

	"github.com/gin-gonic/gin"
)

func TestItemProductHandler(t *testing.T) {
	product := entity.Product{
		Sku:         "test",
		Name:        "name",
		Price:       100,
		ProductType: "dvd",
		Size:        100,
	}

	mockRepo := &MockRepository{}
	mockRepo.On("FindById", "test").Return(&product, nil)

	mockErrorResponse := &MockErrorResponse{}
	mockErrorResponse.On("GenerateResponse", 200, nil)

	sut := &ProductItemHandler{
		Repository:           mockRepo,
		ErrorResponseHandler: mockErrorResponse,
	}

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "test"}}

	sut.GetProduct(c)
}
