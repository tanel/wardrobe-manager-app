package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Redirect redirects to new URL
func Redirect(url string) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		logRequest(r)

		http.Redirect(w, r, url, http.StatusSeeOther)
	}
}
