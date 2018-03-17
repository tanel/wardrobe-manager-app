package controller

import (
	"time"

	"github.com/juju/errors"
	"github.com/satori/go.uuid"
	"github.com/tanel/webapp/db"
	"github.com/tanel/webapp/http"
	"github.com/tanel/webapp/model"
	"golang.org/x/crypto/bcrypt"
)

// PostSignup creates a new user account
func PostSignup(request *http.Request) {
	email := request.FormValue("email")
	if email == "" {
		request.BadRequest("please enter an e-mail")
		return
	}

	password := request.FormValue("password")
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
		b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
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
			PasswordHash: &passwordHash,
		}
		if err := db.InsertUser(request.DB, *user); err != nil {
			request.InternalServerError(errors.Annotate(err, "inserting user into database failed"))
			return
		}
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(password)); err != nil {
			request.Unauthorized("invalid password")
			return
		}
	}

	if ok := request.SetUserID(user.ID); !ok {
		return
	}

	request.Redirect("/")
}
