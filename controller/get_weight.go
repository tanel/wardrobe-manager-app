package controller

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/ui"
	"github.com/tanel/webapp/http"
)

// GetWeight renders an item page
func GetWeight(request *http.Request, userID string) {
	weight, err := db.SelectWeightByID(request.DB, request.ParamByName("id"), userID)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "selecting weight by ID failed"))
		return
	}

	request.Render("weight", ui.NewWeightPage(userID, *weight))
}
