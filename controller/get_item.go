package controller

import (
	"time"

	"github.com/juju/errors"
	"github.com/satori/go.uuid"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/wardrobe-organizer/ui"
	"github.com/tanel/webapp/http"
)

const addToOutfitID = "add-to-outfit-id"

// GetItem renders an item page
func GetItem(request *http.Request, userID string) {
	outfitID, ok := request.SessionValue(addToOutfitID)
	if !ok {
		return
	}

	itemID := request.ParamByName("id")

	item, err := db.SelectItemWithImagesByID(itemID, userID)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "selecting item with images by ID failed"))
		return
	}

	if outfitID != nil {
		outfitItem := model.OutfitItem{}
		outfitItem.ID = uuid.Must(uuid.NewV4()).String()
		outfitItem.OutfitID = *outfitID
		outfitItem.ItemID = itemID
		outfitItem.CreatedAt = time.Now()
		if err := db.InsertOutfitItem(outfitItem); err != nil {
			request.InternalServerError(errors.Annotate(err, "inserting outfit item failed"))
			return
		}

		if ok := request.SetSessionValue(addToOutfitID, ""); !ok {
			return
		}

		request.Redirect("/outfits/" + *outfitID)
		return
	}

	page := ui.NewItemPage(userID, *item)
	request.Render("item", page)
}
