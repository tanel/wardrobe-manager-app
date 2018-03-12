package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/ui"
)

// GetConfirmDeleteWeight renders weight deletion confirmation page
func GetConfirmDeleteWeight(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	weightEntry, err := db.SelectWeightByID(ps.ByName("id"), userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	page := ui.NewWeightPage(userID, *weightEntry)
	if err := Render(w, "confirm-delete-weight", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
