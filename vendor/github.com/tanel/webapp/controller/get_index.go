package controller

import (
	"github.com/tanel/webapp/configuration"
	"github.com/tanel/webapp/http"
	"github.com/tanel/webapp/ui"
)

// GetIndex renders the index page
func GetIndex(request *http.Request) {
	userID, ok := request.UserID()
	if !ok {
		return
	}

	if userID != nil {
		request.Redirect(configuration.LoggedInPage)
		return
	}

	request.Render("index", ui.Page{})
}
