package controller_test

import (
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/tanel/wardrobe-manager-app/controller"
)

func Test_GetIndex_ReturnsResponse_InCaseOfSuccess(t *testing.T) {
	// Arrange
	var ps httprouter.Params
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// Act
	controller.GetIndex(w, req, ps)

	// Assert
	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)
}
