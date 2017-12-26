package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/model"
)

// PostOutfit updates an outfit
func PostOutfit(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	outfit, err := model.NewOutfitForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	outfit.ID = ps.ByName("id")
	outfit.UserID = userID

	if err := db.UpdateOutfit(*outfit); err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/outfits", http.StatusSeeOther)
}
