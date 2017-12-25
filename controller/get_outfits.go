package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/session"
	"github.com/tanel/wardrobe-manager-app/ui"
)

// GetOutfits renders outfits page
func GetOutfits(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	outfits, err := db.SelectOutfitsByUserID(*userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	page := ui.OutfitsPage{
		Page: ui.Page{
			UserID: *userID,
		},
		Outfits: outfits,
	}
	if err := Render(w, "outfits", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
