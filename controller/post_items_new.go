package controller

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/webapp/http"
)

// PostItemsNew creates a new item
func PostItemsNew(request *http.Request, userID string) {
	item, err := model.NewItemForm(request.R())
	if err != nil {
		request.BadRequest(err.Error())
		return
	}

	item.UserID = userID
	if err := db.SaveItem(request.DB, item, userID); err != nil {
		request.InternalServerError(errors.Annotate(err, "saving item failed"))
		return
	}

	request.Redirect(frontPage)
}
