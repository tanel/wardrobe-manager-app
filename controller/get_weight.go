package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/ui"
)

// GetWeight renders an item page
func GetWeight(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	weight, err := db.SelectWeightByID(ps.ByName("id"), userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	page := ui.NewWeightPage(userID, *weight)
	if err := Render(w, "weight", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
