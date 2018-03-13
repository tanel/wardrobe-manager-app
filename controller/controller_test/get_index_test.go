package controller_test

import (
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/tanel/webapp/controller"
	"github.com/tanel/webapp/db"
	"github.com/tanel/webapp/session"
)

func Test_GetIndex_ReturnsResponse_InCaseOfSuccess(t *testing.T) {
	// Arrange
	var ps httprouter.Params
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	databaseConnection := db.Connect("wardrobe", "wardrobe_test")
	sessionStore := session.New("secret", "wardrobe-test")

	// Act
	controller.GetIndex(databaseConnection, sessionStore, w, req, ps)

	// Assert
	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)
}
