package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/wardrobe-organizer/service"
)

// PostItemsNew creates a new item
func PostItemsNew(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	item, err := model.NewItemForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item.UserID = userID

	if err := service.SaveItem(item, userID); err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, frontPage, http.StatusSeeOther)
}
