package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/ui"
)

// GetOutfit renders an outfit page
func GetOutfit(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	outfitID := ps.ByName("id")

	outfit, err := db.SelectOutfitByID(outfitID, userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	outfitItems, err := db.SelectOutfitItemsByOutfitID(outfitID, userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	outfit.OutfitItems = outfitItems

	page := ui.NewOutfitPage(userID, *outfit)
	if err := Render(w, "outfit", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
