package controller

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/webapp/http"
)

// GetRemoveFromOutfit removes an outfit item from outfit
func GetRemoveFromOutfit(request *http.Request, userID string) {
	outfitItemID := request.ParamByName("id")

	outfitID, err := db.SelectOutfitIDByOutfitItemID(outfitItemID, userID)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "selecting outfit ID by outfit item ID failed"))
		return
	}

	if err := db.DeleteOutfitItem(outfitItemID, userID); err != nil {
		request.InternalServerError(errors.Annotate(err, "deleting outfit item failed"))
		return
	}

	request.Redirect("/outfits/" + outfitID)
}
