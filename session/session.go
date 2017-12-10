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
	// Get a session. Get() always returns a session, even if empty.
	session, err := store.Get(r, sessionName)
	if err != nil {
		return nil, errors.Annotate(err, "getting session failed")
	}

	userID, isString := session.Values["user_id"].(string)
	if isString && userID != "" {
		return &userID, nil
	}

	return nil, nil
}

func SetUserID(w http.ResponseWriter, r *http.Request, userID string) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return errors.Annotate(err, "getting session failed")
	}

	// Set some session values.
	session.Values["user_id"] = userID

	if err := session.Save(r, w); err != nil {
		return errors.Annotate(err, "saving session failed")
	}

	return nil
}
