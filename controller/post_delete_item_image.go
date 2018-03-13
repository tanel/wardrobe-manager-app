package controller

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/webapp/session"
)

// PostDeleteItemImage deletes an image
func PostDeleteItemImage(databaseConnection *sql.DB, sessionStore *session.Store, w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	itemImage, err := db.SelectItemImageByID(databaseConnection, ps.ByName("id"), userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if err := db.DeleteItemImage(databaseConnection, ps.ByName("id"), userID); err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/items/"+itemImage.ItemID, http.StatusSeeOther)
}
