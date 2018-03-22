package middleware

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"
	commonhttp "github.com/tanel/webapp/http"
	"github.com/tanel/webapp/session"
)

// RequireUserFunc is a func that requires user ID to execute
type RequireUserFunc func(request *commonhttp.Request, userID string)

// RequireUser wraps regular request to check that user ID is presents in session
func RequireUser(db *sql.DB, sessionStore *session.Store, requireUserFunc RequireUserFunc) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handleRequest(db, sessionStore, w, r, ps, func(request *commonhttp.Request) {
			userID, ok := request.UserID()
			if !ok {
				return
			}

			if userID == nil {
				request.Redirect("/signup")
				return
			}

			requireUserFunc(request, *userID)
		})
	}
}
