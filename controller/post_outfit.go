package controller

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/webapp/http"
)

// PostOutfit updates an outfit
func PostOutfit(request *http.Request, userID string) {
	outfit, err := model.NewOutfitForm(request.R())
	if err != nil {
		request.BadRequest(err.Error())
		return
	}

	outfit.ID = request.ParamByName("id")
	outfit.UserID = userID
	if err := db.UpdateOutfit(request.DB, *outfit); err != nil {
		request.InternalServerError(errors.Annotate(err, "updating outfit in database failed"))
		return
	}

	request.Redirect("/outfits")
}
