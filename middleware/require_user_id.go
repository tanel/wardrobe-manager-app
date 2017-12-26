package middleware

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-manager-app/session"
)

type RequireUserFunc func(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string)

func RequireUser(handlerFunc RequireUserFunc) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		userID, err := session.UserID(r)
		if err != nil {
			log.Println(err)
			http.Error(w, "session error", http.StatusInternalServerError)
			return
		}

		if userID == nil {
			http.Redirect(w, r, "/signup", http.StatusSeeOther)
			return
		}

		handlerFunc(w, r, ps, *userID)
	}
}
