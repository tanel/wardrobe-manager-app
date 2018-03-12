package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-organizer/session"
	"github.com/tanel/wardrobe-organizer/ui"
)

// GetIndex renders the index page
func GetIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, err := session.UserID(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "session error", http.StatusInternalServerError)
		return
	}

	if userID != nil {
		http.Redirect(w, r, "/items", http.StatusSeeOther)
		return
	}

	if err := Render(w, "index", ui.Page{}); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
