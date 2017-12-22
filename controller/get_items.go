package controller

import (
	"log"
	"net/http"

	"github.com/juju/errors"
	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-manager-app/service"
	"github.com/tanel/wardrobe-manager-app/session"
	"github.com/tanel/wardrobe-manager-app/ui"
)

// GetItems renders items page
func GetItems(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	category, err := parseCategory(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "cookie error", http.StatusInternalServerError)
		return
	}

	if err := session.SetCategory(w, r, *category); err != nil {
		log.Println(err)
		http.Error(w, "cookie error", http.StatusInternalServerError)
		return
	}

	categories, err := service.GroupItemsByCategory(*userID, *category)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	page := ui.ItemsPage{
		Page: ui.Page{
			UserID: *userID,
		},
		Categories:       categories,
		SelectedCategory: *category,
	}
	if err := Render(w, "items", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}

func parseCategory(r *http.Request) (*string, error) {
	categoryFromSession, err := session.Category(r)
	if err != nil {
		return nil, errors.Annotate(err, "getting category from session failed")
	}

	categoryFromURL := r.URL.Query().Get("category")
	hasCategoryParameter := false
	for k := range r.URL.Query() {
		if k == "category" {
			hasCategoryParameter = true
			break
		}
	}

	category := ""

	if categoryFromSession != nil {
		category = *categoryFromSession
	}

	if hasCategoryParameter {
		category = categoryFromURL
	}

	return &category, nil
}
