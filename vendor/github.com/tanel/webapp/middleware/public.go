package middleware

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/webapp/session"
)

// PublicFunc is a func that requires no login - its public
type PublicFunc func(databaseConnection *sql.DB, sessionStore *session.Store, w http.ResponseWriter, r *http.Request, ps httprouter.Params)

// HandlePublic wraps regular request
func HandlePublic(databaseConnection *sql.DB, sessionStore *session.Store, handlerFunc PublicFunc) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handlerFunc(databaseConnection, sessionStore, w, r, ps)
	}
}
