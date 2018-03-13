package controller

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/ui"
	"github.com/tanel/webapp/session"
	"github.com/tanel/webapp/template"
)

// GetConfirmDeleteItem renders item deletion confirmation page
func GetConfirmDeleteItem(databaseConnection *sql.DB, sessionStore *session.Store, w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	item, err := db.SelectItemByID(databaseConnection, ps.ByName("id"), userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	page := ui.NewItemPage(userID, *item)
	if err := template.Render(w, "confirm-delete-item", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
