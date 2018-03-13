package controller_test

import (
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/tanel/wardrobe-organizer/controller"
	"github.com/tanel/webapp/db"
)

func Test_GetConfirmDeleteItem_ReturnsResponse_InCaseOfSuccess(t *testing.T) {
	// Arrange
	databaseConnection := db.Connect("wardrobe", "wardrobe_test")
	const itemID = "9924B052-656C-4A99-BE22-5CF0CC2FD287"
	ps := httprouter.Params{
		httprouter.Param{
			Key:   "id",
			Value: itemID,
		},
	}

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	const userID = "3E874918-F54E-49A3-A635-321437A85F00"

	// Act
	controller.GetConfirmDeleteItem(databaseConnection, nil, w, req, ps, userID)

	// Assert
	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)
}
