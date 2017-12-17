package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-manager-app/db"
	//"github.com/tanel/wardrobe-manager-app/model"
	"github.com/tanel/wardrobe-manager-app/session"
	"github.com/tanel/wardrobe-manager-app/ui"
)

func GetItems(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	items, err := db.SelectItemsByUserID(*userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	/* sort group items by category
		var categories model.Category
		for _, item := range items{
	        //if item.Category == ""
		}
	*/

	page := ui.ItemsPage{
		Page: ui.Page{
			UserID: *userID,
		},
		Items: items,
		//Categories: categories,
	}
	if err := Render(w, "items", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
