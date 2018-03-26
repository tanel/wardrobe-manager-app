package controller

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/webapp/http"
)

// GetItemImageThumbnail renders image thumbnail
func GetItemImageThumbnail(request *http.Request, userID string) {
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

	if err := itemImage.LoadThumbnail(); err != nil {
		request.InternalServerError(errors.Annotate(err, "loading thumbnail failed"))
		return
	}

	request.Write(itemImage.Body)
}
