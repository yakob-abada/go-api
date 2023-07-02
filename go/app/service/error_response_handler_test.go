package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yakob-abada/go-api/go/app/model"
)

func TestXxx(t *testing.T) {
	t.Run("InternalServerError", func(t *testing.T) {
		response := &model.ErrorResponse{
			Error: "something wwent wrong",
		}

		sut := NewErrorResponseHandler()
		code, model := sut.GenerateResponse(fmt.Errorf("something wwent wrong"))

		assert.Exactly(t, 500, code)
		assert.Exactly(t, model, response)

	})

	t.Run("UnauthorizedError", func(t *testing.T) {
		response := &model.ErrorResponse{
			Error: "unauthorized error",
		}

		sut := NewErrorResponseHandler()
		code, model := sut.GenerateResponse(NewUnauthorizedError("unauthorized error"))

		assert.Exactly(t, 401, code)
		assert.Exactly(t, model, response)

	})

	t.Run("UnauthorizedError", func(t *testing.T) {
		response := &model.ErrorResponse{
			Error: "unauthorized error",
		}

		sut := NewErrorResponseHandler()
		code, model := sut.GenerateResponse(NewUnauthorizedError("unauthorized error"))

		assert.Exactly(t, 401, code)
		assert.Exactly(t, model, response)

	})

	t.Run("BadRequestError", func(t *testing.T) {
		response := &model.ErrorResponse{
			Error: "badRequest error",
		}

		sut := NewErrorResponseHandler()
		code, model := sut.GenerateResponse(NewBadRequestError("badRequest error"))

		assert.Exactly(t, 400, code)
		assert.Exactly(t, model, response)

	})

	t.Run("UnprocessableEntityError", func(t *testing.T) {
		response := &model.ErrorResponse{
			Error: "unprocessable error",
		}

		sut := NewErrorResponseHandler()
		code, model := sut.GenerateResponse(NewUnprocessableEntityError("unprocessable error"))

		assert.Exactly(t, 422, code)
		assert.Exactly(t, model, response)

	})
}
