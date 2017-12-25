package controller

import (
	"log"
	"net/http"

	"github.com/juju/errors"
	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/session"
)

// GetItemImageThumbnail renders image thumbnail
func GetItemImageThumbnail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	imageID := ps.ByName("id")

	itemImage, err := db.SelectItemImageByID(imageID, *userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	if itemImage == nil {
		http.Error(w, "image not found", http.StatusNotFound)
		return
	}

	if err := itemImage.LoadThumbnail(); err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(itemImage.Body); err != nil {
		log.Println(errors.Annotate(err, "writing image as response failed"))
	}
}