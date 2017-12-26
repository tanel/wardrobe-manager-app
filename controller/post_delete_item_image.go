package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-manager-app/db"
)

// PostDeleteItemImage deletes an image
func PostDeleteItemImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	itemImage, err := db.SelectItemImageByID(ps.ByName("id"), userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if err := db.DeleteItemImage(ps.ByName("id"), userID); err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/items/"+itemImage.ItemID, http.StatusSeeOther)
}
