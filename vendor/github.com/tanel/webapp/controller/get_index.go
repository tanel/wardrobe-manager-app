package controller

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/webapp/session"
	"github.com/tanel/webapp/template"
	"github.com/tanel/webapp/ui"
)

// GetIndex renders the index page
func GetIndex(_ *sql.DB, sessionStore *session.Store, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, err := sessionStore.UserID(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "session error", http.StatusInternalServerError)
		return
	}

	if userID != nil {
		http.Redirect(w, r, "/items", http.StatusSeeOther)
		return
	}

	if err := template.Render(w, "index", ui.Page{}); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
