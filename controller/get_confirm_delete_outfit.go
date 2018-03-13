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

// GetConfirmDeleteOutfit renders outfit deletion confirmation page
func GetConfirmDeleteOutfit(databaseConnection *sql.DB, sessionStore *session.Store, w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	outfit, err := db.SelectOutfitByID(databaseConnection, ps.ByName("id"), userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	page := ui.NewOutfitPage(userID, *outfit)
	if err := template.Render(w, "confirm-delete-outfit", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
