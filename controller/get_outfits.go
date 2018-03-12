package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/ui"
)

// GetOutfits renders outfits page
func GetOutfits(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	outfits, err := db.SelectOutfitsByUserID(userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	page := ui.OutfitsPage{
		Page: ui.Page{
			UserID: userID,
		},
		Outfits: outfits,
	}
	if err := Render(w, "outfits", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
