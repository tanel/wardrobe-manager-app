package controller

import (
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func Test_GetIndex_ReturnsResponse_InCaseOfSuccess(t *testing.T) {
	// Arrange
	var ps httprouter.Params
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	const userID = "3E874918-F54E-49A3-A635-321437A85F00"

	// Act
	GetIndex(w, req, ps, userID)

	// Assert
	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)
}
