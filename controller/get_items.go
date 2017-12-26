package controller

import (
	"log"
	"net/http"

	"github.com/juju/errors"
	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/service"
	"github.com/tanel/wardrobe-manager-app/session"
	"github.com/tanel/wardrobe-manager-app/ui"
)

// GetItems renders items page
func GetItems(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	category, err := parseFilter(r, "category")
	if err != nil {
		log.Println(err)
		http.Error(w, "cookie error", http.StatusInternalServerError)
		return
	}

	if sessionErr := session.SetCategory(w, r, *category); sessionErr != nil {
		log.Println(sessionErr)
		http.Error(w, "cookie error", http.StatusInternalServerError)
		return
	}

	categories, err := service.GroupItemsByCategory(userID, *category)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	brands, err := db.SelectBrandsByUserID(userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	colors, err := db.SelectColorsByUserID(userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	page := ui.ItemsPage{
		Page: ui.Page{
			UserID: userID,
		},
		Categories:       categories,
		SelectedCategory: *category,
		Brands:           brands,
		Colors:           colors,
	}
	if err := Render(w, "items", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}

func parseFilter(r *http.Request, name string) (*string, error) {
	filterFromSession, err := session.Value(r, name)
	if err != nil {
		return nil, errors.Annotatef(err, "getting filter %s from session failed", name)
	}

	filterFromURL := r.URL.Query().Get(name)
	hasFilterParameter := false
	for k := range r.URL.Query() {
		if k == name {
			hasFilterParameter = true
			break
		}
	}

	filter := ""

	if filterFromSession != nil {
		filter = *filterFromSession
	}

	if hasFilterParameter {
		filter = filterFromURL
	}

	return &filter, nil
}
