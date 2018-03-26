package controller

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/webapp/http"
)

// GetItemImage renders iamge
func GetItemImage(request *http.Request, userID string) {
	imageID := request.ParamByName("id")

	itemImage, err := db.SelectItemImageByID(imageID, userID)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "selecting item image by ID failed"))
		return
	}

	if itemImage == nil {
		request.NotFound("image not found")
		return
	}

	if err := itemImage.Load(); err != nil {
		request.InternalServerError(errors.Annotate(err, "loading image failed"))
		return
	}

	request.Write(itemImage.Body)
}
