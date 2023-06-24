package application

import (
	"net/http/httptest"
	"testing"

	"github.com/yakob-abada/go-api/go/app/entity"

	"github.com/gin-gonic/gin"
)

func TestListProductHandler(t *testing.T) {
	products := []entity.Product{
		{
			Sku:         "test",
			Name:        "name",
			Price:       100,
			ProductType: "dvd",
			Size:        100,
		},
	}

	mockRepo := &MockRepository{}
	mockRepo.On("FindAll").Return(&products, nil)

	mockErrorResponse := &MockErrorResponse{}
	mockErrorResponse.On("GenerateResponse", 200, nil)

	sut := &ProductListHandler{
		Repository:           mockRepo,
		ErrorResponseHandler: mockErrorResponse,
	}

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	sut.GetProductList(c)
}
