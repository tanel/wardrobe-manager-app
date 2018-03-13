package controller

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/wardrobe-organizer/ui"
	"github.com/tanel/webapp/session"
	"github.com/tanel/webapp/template"
)

const addToOutfitID = "add-to-outfit-id"

// GetItem renders an item page
func GetItem(databaseConnection *sql.DB, sessionStore *session.Store, w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	outfitID, err := sessionStore.Value(r, addToOutfitID)
	if err != nil {
		log.Println(err)
		http.Error(w, "session error", http.StatusInternalServerError)
		return
	}

	itemID := ps.ByName("id")

	item, err := db.SelectItemWithImagesByID(databaseConnection, itemID, userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	if outfitID != nil {
		outfitItem := model.OutfitItem{}
		outfitItem.ID = uuid.NewV4().String()
		outfitItem.OutfitID = *outfitID
		outfitItem.ItemID = itemID
		outfitItem.CreatedAt = time.Now()
		if err := db.InsertOutfitItem(databaseConnection, outfitItem); err != nil {
			log.Println(err)
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}

		if err := sessionStore.SetValue(w, r, addToOutfitID, ""); err != nil {
			log.Println(err)
			http.Error(w, "session error", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/outfits/"+*outfitID, http.StatusSeeOther)
		return
	}

	page := ui.NewItemPage(userID, *item)
	if err := template.Render(w, "item", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
