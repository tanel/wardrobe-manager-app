package controller

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/webapp/session"
)

// GetLogout logs user out
func GetLogout(_ *sql.DB, sessionStore *session.Store, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := sessionStore.SetUserID(w, r, ""); err != nil {
		log.Println(err)
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
