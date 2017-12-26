package controller

import (
	"log"
	"net/http"

	"github.com/juju/errors"
	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-manager-app/db"
)

// GetItemImage renders iamge
func GetItemImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	imageID := ps.ByName("id")

	itemImage, err := db.SelectItemImageByID(imageID, userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	if itemImage == nil {
		http.Error(w, "image not found", http.StatusNotFound)
		return
	}

	if err := itemImage.Load(); err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(itemImage.Body); err != nil {
		log.Println(errors.Annotate(err, "writing image as response failed"))
	}
}
