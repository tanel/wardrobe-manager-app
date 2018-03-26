package controller

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/ui"
	"github.com/tanel/webapp/http"
)

// GetConfirmDeleteOutfit renders outfit deletion confirmation page
func GetConfirmDeleteOutfit(request *http.Request, userID string) {
	outfit, err := db.SelectOutfitByID(request.ParamByName("id"), userID)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "selecting outfit by ID failed"))
		return
	}

	request.Render("confirm-delete-outfit", ui.NewOutfitPage(userID, *outfit))
}
