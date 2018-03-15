package controller

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/webapp/http"
)

// PostDeleteWeight deletes a weight
func PostDeleteWeight(request *http.Request, userID string) {
	if err := db.DeleteWeight(request.DB, request.ParamByName("id"), userID); err != nil {
		request.InternalServerError(errors.Annotate(err, "deleting weight failed"))
		return
	}

	request.Redirect("/weight")
}
