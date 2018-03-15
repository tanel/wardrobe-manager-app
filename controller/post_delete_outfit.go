package controller

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/webapp/http"
)

// PostDeleteOutfit deletes an outfit
func PostDeleteOutfit(request *http.Request, userID string) {
	if err := db.DeleteOutfit(request.DB, request.ParamByName("id"), userID); err != nil {
		request.InternalServerError(errors.Annotate(err, "deleting outfit failed"))
		return
	}

	request.Redirect("/outfits")
}
