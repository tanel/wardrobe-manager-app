package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-manager-app/ui"
)

// GetIndex renders the index page
func GetIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	if err := Render(w, "index", ui.Page{}); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
