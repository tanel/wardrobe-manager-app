package controller

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/webapp/session"
)

// PostNewOutfit saves a new outfit into database
func PostNewOutfit(databaseConnection *sql.DB, sessionStore *session.Store, w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	outfit, err := model.NewOutfitForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	outfit.ID = uuid.NewV4().String()
	outfit.UserID = userID
	outfit.CreatedAt = time.Now()

	if err := db.InsertOutfit(databaseConnection, *outfit); err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/outfits/"+outfit.ID, http.StatusSeeOther)
}
