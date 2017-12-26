package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-manager-app/db"
)

// PostDeleteItem deletes an item
func PostDeleteItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	if err := db.DeleteItem(ps.ByName("id"), userID); err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, frontPage, http.StatusSeeOther)
}
