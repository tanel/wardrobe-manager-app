package controller

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/webapp/http"
)

// PostDeleteItem deletes an item
func PostDeleteItem(request *http.Request, userID string) {
	if err := db.DeleteItem(request.ParamByName("id"), userID); err != nil {
		request.InternalServerError(errors.Annotate(err, "deleting item failed"))
		return
	}

	request.Redirect("/items")
}
