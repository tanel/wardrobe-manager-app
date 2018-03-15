package controller

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/webapp/http"
)

// PostWeight updates a weight entry
func PostWeight(request *http.Request, userID string) {
	weightEntry, err := model.NewWeightEntryForm(request.R())
	if err != nil {
		request.BadRequest(err.Error())
		return
	}

	weightEntry.ID = request.ParamByName("id")
	weightEntry.UserID = userID
	if err := db.UpdateWeight(request.DB, *weightEntry); err != nil {
		request.InternalServerError(errors.Annotate(err, "updating weight failed"))
		return
	}

	request.Redirect("/weight")
}
