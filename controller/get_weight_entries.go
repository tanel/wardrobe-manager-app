package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/ui"
)

// GetWeightEntries renders weight entries page
func GetWeightEntries(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	weights, err := db.SelectWeightsByUserID(userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	page, err := ui.NewWeightEntriesPage(userID, weights)
	if err != nil {
		log.Println(err)
		http.Error(w, "page error", http.StatusInternalServerError)
		return
	}

	if err := Render(w, "weight-entries", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
