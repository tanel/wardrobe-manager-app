package session

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/juju/errors"
)

const sessionName = "wardrobe-app-session"

// FIXME: get secret from environment
var store = sessions.NewCookieStore([]byte("C93B74DA-4D85-418C-B513-3BEDE6BFCECC"))

func UserID(r *http.Request) (*string, error) {
	return Value(r, "user_id")
}

func Category(r *http.Request) (*string, error) {
	return Value(r, "category")
}

func Value(r *http.Request, key string) (*string, error) {
	// Get a session. Get() always returns a session, even if empty.
	session, err := store.Get(r, sessionName)
	if err != nil {
		return nil, errors.Annotate(err, "getting session failed")
	}

	value, isString := session.Values[key].(string)
	if isString && value != "" {
		return &value, nil
	}

	return nil, nil
}

func SetUserID(w http.ResponseWriter, r *http.Request, userID string) error {
	return SetValue(w, r, "user_id", userID)
}

func SetCategory(w http.ResponseWriter, r *http.Request, category string) error {
	return SetValue(w, r, "category", category)
}

func SetValue(w http.ResponseWriter, r *http.Request, key string, value string) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return errors.Annotate(err, "getting session failed")
	}

	// Set some session values.
	session.Values[key] = value

	if err := session.Save(r, w); err != nil {
		return errors.Annotate(err, "saving session failed")
	}

	return nil
}
