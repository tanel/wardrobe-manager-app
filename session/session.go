package session

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/juju/errors"
)

const (
	sessionName         = "wardrobe-app-session"
	unsecureCredentials = "C93B74DA-4D85-418C-B513-3BEDE6BFCECC"

	// AddToOutfitID is session key for adding items to an outfit ID
	AddToOutfitID = "add-to-outfit-id"
)

var store *sessions.CookieStore

func init() {
	sessionSecret := os.Getenv("SESSION_SECRET")
	if sessionSecret == "" {
		sessionSecret = unsecureCredentials
		log.Println("Warning: SESSION_SECRET not set in env")
	}

	store = sessions.NewCookieStore([]byte(sessionSecret))
}

// UserID returns user ID from session
func UserID(r *http.Request) (*string, error) {
	return Value(r, "user_id")
}

// Value returns reads a value from session by key
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

// SetUserID sets user ID in session
func SetUserID(w http.ResponseWriter, r *http.Request, userID string) error {
	return SetValue(w, r, "user_id", userID)
}

// SetValue sets a value in session
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
