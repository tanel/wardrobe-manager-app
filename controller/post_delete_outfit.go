package controller

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/webapp/session"
)

// PostDeleteOutfit deletes an outfit
func PostDeleteOutfit(databaseConnection *sql.DB, sessionStore *session.Store, w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	if err := db.DeleteOutfit(databaseConnection, ps.ByName("id"), userID); err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/outfits", http.StatusSeeOther)
}
