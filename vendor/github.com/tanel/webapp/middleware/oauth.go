package middleware

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/webapp/configuration"
	commonhttp "github.com/tanel/webapp/http"
	"github.com/tanel/webapp/session"
)

// OAuth2Func is a func that handles OAuth2 requests
type OAuth2Func func(r *commonhttp.Request, cfg configuration.OAuth2)

// HandleOAuth2 wraps regular request
func HandleOAuth2(db *sql.DB, sessionStore *session.Store, cfg configuration.OAuth2, handlerFunc OAuth2Func) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		logRequest(r)

		request, err := commonhttp.NewRequest(db, sessionStore, w, r, ps)
		if err != nil {
			log.Println(err)
			http.Error(w, "handling request failed", http.StatusInternalServerError)
			return
		}

		handlerFunc(request, cfg)
	}
}
