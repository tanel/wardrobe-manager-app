package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-organizer/db"
)

// GetRemoveFromOutfit removes an outfit item from outfit
func GetRemoveFromOutfit(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	outfitItemID := ps.ByName("id")

	outfitID, err := db.SelectOutfitIDByOutfitItemID(outfitItemID, userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	if err := db.DeleteOutfitItem(outfitItemID, userID); err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/outfits/"+outfitID, http.StatusSeeOther)

}
