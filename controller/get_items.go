package controller

import (
	"log"
	"net/http"

	"github.com/juju/errors"
	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/model"
	"github.com/tanel/wardrobe-manager-app/service"
	"github.com/tanel/wardrobe-manager-app/session"
	"github.com/tanel/wardrobe-manager-app/ui"
)

// GetItems renders items page
func GetItems(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	f, err := newFilters(w, r)
	if err != nil {
		log.Println(err)
		http.Error(w, "filters error", http.StatusInternalServerError)
		return
	}

	var outfit *model.Outfit
	if f.outfitID != "" {
		outfit, err = db.SelectOutfitByID(f.outfitID, userID)
		if err != nil {
			log.Println(err)
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}
	}

	itemCategories, err := service.GroupItemsByCategory(userID, f.category, f.brand, f.color)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	categories, err := db.SelectCategoriesByUserID(userID)
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
		ItemCategories:   itemCategories,
		Categories:       categories,
		Brands:           brands,
		Colors:           colors,
		SelectedOutfit:   outfit,
		SelectedCategory: f.category,
		SelectedBrand:    f.brand,
		SelectedColor:    f.color,
	}
	if err := Render(w, "items", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}

func parseFilter(r *http.Request, name string) (string, error) {
	filterFromSession, err := session.Value(r, name)
	if err != nil {
		return "", errors.Annotatef(err, "getting filter %s from session failed", name)
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

	return filter, nil
}

func handleParam(w http.ResponseWriter, r *http.Request, name string) (string, error) {
	value, err := parseFilter(r, name)
	if err != nil {
		return "", errors.Annotatef(err, "reading %s from cookie failed", name)
	}

	if sessionErr := session.SetValue(w, r, name, value); sessionErr != nil {
		return "", errors.Annotatef(err, "setting %s in session failed", name)
	}

	return value, nil
}

type filters struct {
	category string
	brand    string
	color    string
	outfitID string
}

func newFilters(w http.ResponseWriter, r *http.Request) (*filters, error) {
	var result filters
	var err error

	result.category, err = handleParam(w, r, "category")
	if err != nil {
		return nil, errors.Annotate(err, "handling category param failed")
	}

	result.brand, err = handleParam(w, r, "brand")
	if err != nil {
		return nil, errors.Annotate(err, "handling brand param failed")
	}

	result.color, err = handleParam(w, r, "color")
	if err != nil {
		return nil, errors.Annotate(err, "handling color param failed")
	}

	result.outfitID, err = handleParam(w, r, session.AddToOutfitID)
	if err != nil {
		return nil, errors.Annotate(err, "handling outfit ID param failed")
	}

	return &result, nil
}
