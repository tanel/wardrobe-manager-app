package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/ui"
)

// GetConfirmDeleteOutfit renders outfit deletion confirmation page
func GetConfirmDeleteOutfit(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	outfit, err := db.SelectOutfitByID(ps.ByName("id"), userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	page := ui.NewOutfitPage(userID, *outfit)
	if err := Render(w, "confirm-delete-outfit", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
