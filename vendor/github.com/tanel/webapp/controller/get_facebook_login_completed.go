package controller

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/juju/errors"
	"github.com/satori/go.uuid"
	"github.com/tanel/webapp/configuration"
	"github.com/tanel/webapp/db"
	"github.com/tanel/webapp/http"
	"github.com/tanel/webapp/model"
)

// GetFacebookLoginCompleted finishes FB signup/login
func GetFacebookLoginCompleted(request *http.Request, cfg configuration.OAuth2) {
	code := request.QueryParamByName("code")
	fbUser, err := facebookUser(code, cfg)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "getting Facebook user failed"))
		return
	}

	log.Println("FB login completed", fbUser)

	// select user by user.Email
	user, err := db.SelectUserByEmail(request.DB, fbUser.Email)
	if err != nil {
		request.InternalServerError(errors.Annotatef(err, "selecting user by email failed, email=%s", fbUser.Email))
		return
	}

	if user == nil {
		log.Println(fbUser.Email, "not found, creating new user")

		user = &model.User{
			Base: model.Base{
				ID:        uuid.Must(uuid.NewV4()).String(),
				CreatedAt: time.Now(),
			},
			Email:   fbUser.Email,
			Name:    fbUser.Name,
			Picture: fbUser.Picture,
		}
		if err := db.InsertUser(request.DB, *user); err != nil {
			request.InternalServerError(errors.Annotate(err, "inserting user failed"))
			return
		}
	} else {
		// Update user data from Facebook
		user.Picture = fbUser.Picture
		if err := db.UpdateUser(request.DB, *user); err != nil {
			request.InternalServerError(errors.Annotate(err, "updating user failed"))
			return
		}
	}

	log.Println("logging user in, ID=", user.ID)

	if ok := request.SetUserID(user.ID); !ok {
		return
	}

	request.Redirect(startPage)
}

// private methods

func facebookUser(code string, cfg configuration.OAuth2) (*model.User, error) {
	// FIXME: load configuration only once
	conf := cfg.Facebook()
	ctx := context.Background()

	token, err := conf.Exchange(ctx, code)
	if err != nil {
		return nil, errors.Annotate(err, "exchanging oauth token failed")
	}

	url := "https://graph.facebook.com/me?fields=name,email,picture.type(large)"
	client := conf.Client(ctx, token)
	resp, err := client.Get(url)
	if err != nil {
		return nil, errors.Annotate(err, "getting user data from FB failed")
	}

	var profile map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, errors.Annotate(err, "decoding oauth2 user data failed")
	}

	var user model.User
	user.Email = profile["email"].(string)
	user.Name = profile["name"].(string)
	if pic, isMap := profile["picture"].(map[string]interface{}); isMap {
		if data, isMap := pic["data"].(map[string]interface{}); isMap {
			if s, isString := data["url"].(string); isString {
				user.Picture = &s
			}
		}
	}

	return &user, nil
}
