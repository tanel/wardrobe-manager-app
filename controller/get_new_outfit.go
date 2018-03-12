package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/wardrobe-organizer/ui"
)

// GetNewOutfit renders new outfit page
func GetNewOutfit(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	page := ui.NewOutfitPage(userID, model.Outfit{})
	if err := Render(w, "new-outfit", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
