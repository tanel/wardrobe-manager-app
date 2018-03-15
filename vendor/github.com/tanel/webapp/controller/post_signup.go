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
	var user model.User
	user.Email = request.FormValue("email")
	if user.Email == "" {
		request.BadRequest("please enter an e-mail")
		return
	}

	password := request.FormValue("password")
	if password == "" {
		request.BadRequest("please enter a password")
		return
	}

	if err := db.SelectUserByEmail(request.DB, user.Email, &user); err != nil {
		request.InternalServerError(errors.Annotate(err, "selecting user by email failed"))
		return
	}

	if user.ID == "" {
		user.ID = uuid.Must(uuid.NewV4()).String()
		user.CreatedAt = time.Now()

		b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			request.InternalServerError(errors.Annotate(err, "hashing password failed"))
			return
		}

		user.PasswordHash = string(b)

		if err := db.InsertUser(request.DB, user); err != nil {
			request.InternalServerError(errors.Annotate(err, "inserting user into database failed"))
			return
		}
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
			request.Unauthorized("invalid password")
			return
		}
	}

	if ok := request.SetUserID(user.ID); !ok {
		return
	}

	request.Redirect("/")
}
