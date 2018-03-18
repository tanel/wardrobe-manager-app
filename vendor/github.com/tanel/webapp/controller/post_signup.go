package controller

import (
	"strings"

	"github.com/juju/errors"
	"github.com/tanel/webapp/configuration"
	"github.com/tanel/webapp/db"
	commonemail "github.com/tanel/webapp/email"
	"github.com/tanel/webapp/http"
)

// PostSignup creates a new user account
func PostSignup(request *http.Request) {
	email := strings.TrimSpace(request.FormValue("email"))
	if email == "" {
		request.BadRequest("please enter an e-mail")
		return
	}

	password := strings.TrimSpace(request.FormValue("password"))
	if password == "" {
		request.BadRequest("please enter a password")
		return
	}

	user, err := db.SelectUserByEmail(request.DB, email)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "selecting user by email failed"))
		return
	}

	if user == nil {
		created, ok := request.CreateUser(email, password)
		if !ok {
			return
		}

		user = created

	} else if user.PasswordHash != nil {
		if err := user.CheckPassword(password); err != nil {
			request.Unauthorized("invalid password")
			return
		}

	} else {
		if err := commonemail.SendSetPasswordLink(user.Email); err != nil {
			request.InternalServerError(errors.Annotate(err, "sending set password link failed"))
			return
		}

		request.Redirect("/set-password-link-sent")
	}

	if ok := request.SetUserID(user.ID); !ok {
		return
	}

	request.Redirect(configuration.LoggedInPage)
}
