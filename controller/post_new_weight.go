package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/model"
	"github.com/tanel/wardrobe-manager-app/session"
)

// PostNewWeight saves a new weight into database
func PostNewWeight(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	weightEntry, err := model.NewWeightEntryForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	weightEntry.ID = uuid.NewV4().String()
	weightEntry.UserID = *userID

	if err := db.InsertWeight(*weightEntry); err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/weight", http.StatusSeeOther)
}
