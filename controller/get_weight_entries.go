package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/session"
	"github.com/tanel/wardrobe-manager-app/ui"
)

// GetWeightEntries renders weight entries page
func GetWeightEntries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	weights, err := db.SelectWeightsByUserID(*userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	page := ui.NewWeightEntriesPage(*userID, weights)
	if err := Render(w, "weight-entries", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
