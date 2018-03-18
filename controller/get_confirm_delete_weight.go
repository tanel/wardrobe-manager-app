package controller

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/ui"
	"github.com/tanel/webapp/http"
)

// GetConfirmDeleteWeight renders weight deletion confirmation page
func GetConfirmDeleteWeight(request *http.Request, userID string) {
	weightEntry, err := db.SelectWeightByID(request.DB, request.ParamByName("id"), userID)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "selecting weight by ID failed"))
		return
	}

	request.Render("confirm-delete-weight", ui.NewWeightPage(userID, *weightEntry))
}
