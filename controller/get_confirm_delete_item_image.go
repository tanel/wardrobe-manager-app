package controller

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/ui"
	"github.com/tanel/webapp/http"
)

// GetConfirmDeleteItemImage renders image deletion confirmation page
func GetConfirmDeleteItemImage(request *http.Request, userID string) {
	itemImage, err := db.SelectItemImageByID(request.ParamByName("id"), userID)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "selecting item image by ID failed"))
		return
	}

	request.Render("confirm-delete-item-image", ui.NewItemImagePage(userID, *itemImage))
}
