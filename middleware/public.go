package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// PublicFunc is a func that requires no login - its public
type PublicFunc func(w http.ResponseWriter, r *http.Request, ps httprouter.Params)

// Public wraps regular request
func Public(handlerFunc PublicFunc) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handlerFunc(w, r, ps)
	}
}
