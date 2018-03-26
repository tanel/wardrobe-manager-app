package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	commonhttp "github.com/tanel/webapp/http"
)

// PublicFunc is a func that requires no login - its public
type PublicFunc func(r *commonhttp.Request)

// HandlePublic wraps regular request
func HandlePublic(publicFunc PublicFunc) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handleRequest(w, r, ps, func(request *commonhttp.Request) {
			publicFunc(request)
		})
	}
}
