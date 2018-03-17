package middleware

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	commonhttp "github.com/tanel/webapp/http"
	"github.com/tanel/webapp/session"
)

// PublicFunc is a func that requires no login - its public
type PublicFunc func(r *commonhttp.Request)

// HandlePublic wraps regular request
func HandlePublic(db *sql.DB, sessionStore *session.Store, handlerFunc PublicFunc) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		logRequest(r)

		request, err := commonhttp.NewRequest(db, sessionStore, w, r, ps)
		if err != nil {
			log.Println(err)
			http.Error(w, "handling request failed", http.StatusInternalServerError)
			return
		}

		handlerFunc(request)
	}
}
