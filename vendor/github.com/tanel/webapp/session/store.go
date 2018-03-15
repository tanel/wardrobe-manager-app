package session

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/juju/errors"
)

// Store wraps actual session store logic
type Store struct {
	sessionName string
	cookieStore *sessions.CookieStore
}

// New returns new session store
func New(sessionSecret, sessionName string) *Store {
	if sessionSecret == "" {
		log.Panic("session secret it mandatory")
	}

	if sessionName == "" {
		log.Panic("session name is mandatory")
	}

	return &Store{
		cookieStore: sessions.NewCookieStore([]byte(sessionSecret)),
		sessionName: sessionName,
	}
}

// UserID returns user ID from session
func (store *Store) UserID(r *http.Request) (*string, error) {
	return store.Value(r, "user_id")
}

// SetUserID sets user ID in session
func (store *Store) SetUserID(w http.ResponseWriter, r *http.Request, userID string) error {
	return store.SetValue(w, r, "user_id", userID)
}

// Value returns reads a value from session by key
func (store *Store) Value(r *http.Request, key string) (*string, error) {
	// Get a session. Get() always returns a session, even if empty.
	session, err := store.cookieStore.Get(r, store.sessionName)
	if err != nil {
		return nil, errors.Annotate(err, "getting session failed")
	}

	value, isString := session.Values[key].(string)
	if isString && value != "" {
		return &value, nil
	}

	return nil, nil
}

// SetValue sets a value in session
func (store *Store) SetValue(w http.ResponseWriter, r *http.Request, key string, value string) error {
	session, err := store.cookieStore.Get(r, store.sessionName)
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
