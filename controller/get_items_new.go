package controller

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/wardrobe-organizer/ui"
	"github.com/tanel/webapp/session"
	"github.com/tanel/webapp/template"
)

// GetItemsNew renders new item page
func GetItemsNew(db *sql.DB, sessionStore *session.Store, w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	category := r.URL.Query().Get("category")

	page := ui.NewItemPage(userID, model.Item{
		Currency: "EUR",
		Quantity: 1,
		Category: category,
	})
	if err := template.Render(w, "items-new", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
