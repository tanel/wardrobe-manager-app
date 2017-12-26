package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/ui"
)

// GetConfirmDeleteItemImage renders image deletion confirmation page
func GetConfirmDeleteItemImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	itemImage, err := db.SelectItemImageByID(ps.ByName("id"), userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	page := ui.NewItemImagePage(userID, *itemImage)
	if err := Render(w, "confirm-delete-item-image", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
