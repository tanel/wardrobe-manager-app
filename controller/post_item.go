package controller

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/webapp/http"
)

// PostItem updates an item
func PostItem(request *http.Request, userID string) {
	item, err := model.NewItemForm(request.R())
	if err != nil {
		request.BadRequest(err.Error())
		return
	}

	item.ID = request.ParamByName("id")
	if err := db.SaveItem(request.DB, item, userID); err != nil {
		request.InternalServerError(errors.Annotate(err, "saving item failed"))
		return
	}

	request.Redirect("/items")
}
