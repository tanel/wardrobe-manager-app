package controller

import (
	"time"

	"github.com/juju/errors"
	"github.com/satori/go.uuid"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/webapp/http"
)

// PostNewOutfit saves a new outfit into database
func PostNewOutfit(request *http.Request, userID string) {
	outfit, err := model.NewOutfitForm(request.R())
	if err != nil {
		request.BadRequest(err.Error())
		return
	}

	outfit.ID = uuid.Must(uuid.NewV4()).String()
	outfit.UserID = userID
	outfit.CreatedAt = time.Now()
	if err := db.InsertOutfit(request.DB, *outfit); err != nil {
		request.InternalServerError(errors.Annotate(err, "inserting outfit into database failed"))
		return
	}

	request.Redirect("/outfits/" + outfit.ID)
}
