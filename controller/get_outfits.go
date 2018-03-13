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
	commonui "github.com/tanel/webapp/ui"
)

// GetOutfits renders outfits page
func GetOutfits(databaseConnection *sql.DB, sessionStore *session.Store, w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	outfits, err := db.SelectOutfitsByUserID(databaseConnection, userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	page := ui.OutfitsPage{
		Page: commonui.Page{
			UserID: userID,
		},
		Outfits: outfits,
	}
	if err := template.Render(w, "outfits", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
