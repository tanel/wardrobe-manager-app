package middleware

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"
	commonhttp "github.com/tanel/webapp/http"
	"github.com/tanel/webapp/session"
)

// PublicFunc is a func that requires no login - its public
type PublicFunc func(r *commonhttp.Request)

// HandlePublic wraps regular request
func HandlePublic(db *sql.DB, sessionStore *session.Store, publicFunc PublicFunc) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handleRequest(db, sessionStore, w, r, ps, func(request *commonhttp.Request) {
			publicFunc(request)
		})
	}
}
