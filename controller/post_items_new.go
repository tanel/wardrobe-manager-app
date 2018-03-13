package controller

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/wardrobe-organizer/service"
	"github.com/tanel/webapp/session"
)

// PostItemsNew creates a new item
func PostItemsNew(databaseConnection *sql.DB, sessionStore *session.Store, w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	item, err := model.NewItemForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item.UserID = userID

	if err := service.SaveItem(databaseConnection, item, userID); err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, frontPage, http.StatusSeeOther)
}
