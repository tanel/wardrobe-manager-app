package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-organizer/session"
)

// GetLogout logs user out
func GetLogout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := session.SetUserID(w, r, ""); err != nil {
		log.Println(err)
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
