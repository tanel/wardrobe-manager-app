package controller

import (
	"github.com/tanel/webapp/http"
)

// GetSignup renders signup page
func GetSignup(request *http.Request) {
	userID, ok := request.UserID()
	if !ok {
		return
	}

	if userID != nil {
		request.Redirect("/")
		return
	}

	request.Render("signup", nil)
}
