package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/model"
	"github.com/tanel/wardrobe-manager-app/session"
	"github.com/tanel/wardrobe-manager-app/ui"
)

// GetItem renders an item page
func GetItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	outfitID, err := session.Value(r, session.AddToOutfitID)
	if err != nil {
		log.Println(err)
		http.Error(w, "cookie error", http.StatusInternalServerError)
		return
	}

	itemID := ps.ByName("id")

	item, err := db.SelectItemWithImagesByID(itemID, userID)
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
		if err := db.InsertOutfitItem(outfitItem); err != nil {
			log.Println(err)
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}

		session.SetValue(w, r, session.AddToOutfitID, "")

		http.Redirect(w, r, "/outfits/"+*outfitID, http.StatusSeeOther)
		return
	}

	page := ui.NewItemPage(userID, *item)
	if err := Render(w, "item", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
