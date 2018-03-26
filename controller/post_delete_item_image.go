package controller

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/webapp/http"
)

// PostDeleteItemImage deletes an image
func PostDeleteItemImage(request *http.Request, userID string) {
	itemImage, err := db.SelectItemImageByID(request.ParamByName("id"), userID)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "selecting item image by ID failed"))
		return
	}

	if err := db.DeleteItemImage(request.ParamByName("id"), userID); err != nil {
		request.InternalServerError(errors.Annotate(err, "deleting item image failed"))
		return
	}

	request.Redirect("/items/" + itemImage.ItemID)
}
