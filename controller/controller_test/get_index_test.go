package controller_test

import (
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/tanel/webapp/controller"
	"github.com/tanel/webapp/db"
	commonhttp "github.com/tanel/webapp/http"
)

func Test_GetIndex_ReturnsResponse_InCaseOfSuccess(t *testing.T) {
	// Arrange
	db.Init("wardrobe_test")
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	var ps httprouter.Params
	request, err := commonhttp.NewRequest(w, r, ps)

	// Act
	controller.GetIndex(request)

	// Assert
	assert.NoError(t, err)
	res := w.Result()
	assert.Equal(t, 200, res.StatusCode)
}
