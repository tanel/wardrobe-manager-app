package middleware

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	commonhttp "github.com/tanel/webapp/http"
	"github.com/tanel/webapp/session"
)

// RequestHandlerFunc is a request handler
type RequestHandlerFunc func(request *commonhttp.Request)

func handleRequest(db *sql.DB, sessionStore *session.Store, w http.ResponseWriter, r *http.Request, ps httprouter.Params, requestHandlerFunc RequestHandlerFunc) {
	logRequest(r)

	request, err := commonhttp.NewRequest(db, sessionStore, w, r, ps)
	if err != nil {
		log.Println(err)
		http.Error(w, "handling request failed", http.StatusInternalServerError)
		return
	}

	requestHandlerFunc(request)
}
