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

// GetNewOutfit renders new outfit page
func GetNewOutfit(databaseConnection *sql.DB, sessionStore *session.Store, w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	page := ui.NewOutfitPage(userID, model.Outfit{})
	if err := template.Render(w, "new-outfit", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
