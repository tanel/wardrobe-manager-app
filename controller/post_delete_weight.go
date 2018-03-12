package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-organizer/db"
)

// PostDeleteWeight deletes a weight
func PostDeleteWeight(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	if err := db.DeleteWeight(ps.ByName("id"), userID); err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/weight", http.StatusSeeOther)
}
