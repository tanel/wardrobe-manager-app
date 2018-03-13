package controller

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/ui"
	"github.com/tanel/webapp/session"
	"github.com/tanel/webapp/template"
)

// GetWeightEntries renders weight entries page
func GetWeightEntries(databaseConnection *sql.DB, sessionStore *session.Store, w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	weights, err := db.SelectWeightsByUserID(databaseConnection, userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	page, err := ui.NewWeightEntriesPage(userID, weights)
	if err != nil {
		log.Println(err)
		http.Error(w, "page error", http.StatusInternalServerError)
		return
	}

	if err := template.Render(w, "weight-entries", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
