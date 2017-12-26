package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/model"
)

// PostWeight updates a weight entry
func PostWeight(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	weightEntry, err := model.NewWeightEntryForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	weightEntry.ID = ps.ByName("id")
	weightEntry.UserID = userID

	if err := db.UpdateWeight(*weightEntry); err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/weight", http.StatusSeeOther)
}
