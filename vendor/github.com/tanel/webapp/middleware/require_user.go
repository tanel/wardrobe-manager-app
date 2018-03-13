package middleware

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/webapp/session"
)

// RequireUserFunc is a func that requires user ID to execute
type RequireUserFunc func(databaseConnection *sql.DB, sessionStore *session.Store, w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string)

// RequireUser wraps regular request to check that user ID is presents in session
func RequireUser(databaseConnection *sql.DB, sessionStore *session.Store, handlerFunc RequireUserFunc) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		userID, err := sessionStore.UserID(r)
		if err != nil {
			log.Println(err)
			http.Error(w, "session error", http.StatusInternalServerError)
			return
		}

		if userID == nil {
			http.Redirect(w, r, "/signup", http.StatusSeeOther)
			return
		}

		handlerFunc(databaseConnection, sessionStore, w, r, ps, *userID)
	}
}
