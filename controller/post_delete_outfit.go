package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-manager-app/db"
)

// PostDeleteOutfit deletes an outfit
func PostDeleteOutfit(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	if err := db.DeleteOutfit(ps.ByName("id"), userID); err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/outfits", http.StatusSeeOther)
}
