package controller

import (
	"time"

	"github.com/juju/errors"
	"github.com/satori/go.uuid"
	"github.com/tanel/webapp/db"
	"github.com/tanel/webapp/email"
	"github.com/tanel/webapp/http"
	"github.com/tanel/webapp/model"
	"golang.org/x/crypto/bcrypt"
)

// PostSignup creates a new user account
func PostSignup(request *http.Request) {
	signupEmail := request.FormValue("email")
	if signupEmail == "" {
		request.BadRequest("please enter an e-mail")
		return
	}

	signupPassword := request.FormValue("password")
	if signupPassword == "" {
		request.BadRequest("please enter a password")
		return
	}

	user, err := db.SelectUserByEmail(request.DB, signupEmail)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "selecting user by email failed"))
		return
	}

	if user == nil {
		b, err := bcrypt.GenerateFromPassword([]byte(signupPassword), bcrypt.DefaultCost)
		if err != nil {
			request.InternalServerError(errors.Annotate(err, "hashing password failed"))
			return
		}

		passwordHash := string(b)
		user = &model.User{
			Base: model.Base{
				ID:        uuid.Must(uuid.NewV4()).String(),
				CreatedAt: time.Now(),
			},
			Name:         signupEmail,
			Email:        signupEmail,
			PasswordHash: &passwordHash,
		}
		if err := db.InsertUser(request.DB, *user); err != nil {
			request.InternalServerError(errors.Annotate(err, "inserting user into database failed"))
			return
		}

	} else if user.PasswordHash != nil {
		if err := bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(signupPassword)); err != nil {
			request.Unauthorized("invalid password")
			return
		}

	} else {
		if err := email.SendSetPasswordLink(user.Email); err != nil {
			request.InternalServerError(errors.Annotate(err, "sending set password link failed"))
			return
		}

		request.Redirect("/set-password-link-sent")
	}

	if ok := request.SetUserID(user.ID); !ok {
		return
	}

	request.Redirect(startPage)
}
