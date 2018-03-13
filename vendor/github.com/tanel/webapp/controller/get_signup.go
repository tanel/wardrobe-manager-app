package controller

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/webapp/session"
	"github.com/tanel/webapp/template"
)

// GetSignup renders signup page
func GetSignup(_ *sql.DB, sessionStore *session.Store, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, err := sessionStore.UserID(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "session error", http.StatusInternalServerError)
		return
	}

	if userID != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if err := template.Render(w, "signup", nil); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
