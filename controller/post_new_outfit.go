package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/model"
	"github.com/tanel/wardrobe-manager-app/session"
)

// PostNewOutfit saves a new outfit into database
func PostNewOutfit(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, err := session.UserID(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "session error", http.StatusInternalServerError)
		return
	}

	if userID == nil {
		http.Redirect(w, r, loginPage, http.StatusSeeOther)
		return
	}

	outfit, err := model.NewOutfitForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	outfit.ID = uuid.NewV4().String()
	outfit.UserID = *userID
	outfit.CreatedAt = time.Now()

	if err := db.InsertOutfit(*outfit); err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/outfits", http.StatusSeeOther)
}