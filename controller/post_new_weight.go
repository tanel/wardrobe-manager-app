package controller

import (
	"github.com/juju/errors"
	"github.com/satori/go.uuid"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/webapp/http"
)

// PostNewWeight saves a new weight into database
func PostNewWeight(request *http.Request, userID string) {
	weightEntry, err := model.NewWeightEntryForm(request.R())
	if err != nil {
		request.BadRequest(err.Error())
		return
	}

	weightEntry.ID = uuid.Must(uuid.NewV4()).String()
	weightEntry.UserID = userID
	if err := db.InsertWeight(*weightEntry); err != nil {
		request.InternalServerError(errors.Annotate(err, "inserting weight into database failed"))
		return
	}

	request.Redirect("/weight")
}
