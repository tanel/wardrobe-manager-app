package controller

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/ui"
	"github.com/tanel/webapp/http"
)

// GetConfirmDeleteItem renders item deletion confirmation page
func GetConfirmDeleteItem(request *http.Request, userID string) {
	item, err := db.SelectItemByID(request.DB, request.ParamByName("id"), userID)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "selecting item by ID failed"))
		return
	}

	request.Render("confirm-delete-item", ui.NewItemPage(userID, *item))
}
