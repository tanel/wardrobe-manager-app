package controller

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/ui"
	"github.com/tanel/webapp/http"
)

// GetWeightEntries renders weight entries page
func GetWeightEntries(request *http.Request, userID string) {
	weights, err := db.SelectWeightsByUserID(userID)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "selecting weights by user ID failed"))
		return
	}

	page, err := ui.NewWeightEntriesPage(userID, weights)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "initializing weight entries page failed"))
		return
	}

	request.Render("weight-entries", page)
}
