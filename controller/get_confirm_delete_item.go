package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/ui"
)

// GetConfirmDeleteItem renders item deletion confirmation page
func GetConfirmDeleteItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	item, err := db.SelectItemByID(ps.ByName("id"), userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	page := ui.NewItemPage(userID, *item)
	if err := Render(w, "confirm-delete-item", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
