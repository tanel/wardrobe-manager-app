package controller

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/webapp/session"
)

// GetRemoveFromOutfit removes an outfit item from outfit
func GetRemoveFromOutfit(databaseConnection *sql.DB, sessionStore *session.Store, w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	outfitItemID := ps.ByName("id")

	outfitID, err := db.SelectOutfitIDByOutfitItemID(databaseConnection, outfitItemID, userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	if err := db.DeleteOutfitItem(databaseConnection, outfitItemID, userID); err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/outfits/"+outfitID, http.StatusSeeOther)

}
