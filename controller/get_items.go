package controller

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/juju/errors"
	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/wardrobe-organizer/service"
	"github.com/tanel/wardrobe-organizer/ui"
	"github.com/tanel/webapp/session"
	"github.com/tanel/webapp/template"
	commonui "github.com/tanel/webapp/ui"
)

// GetItems renders items page
func GetItems(databaseConnection *sql.DB, sessionStore *session.Store, w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	f, err := newFilters(sessionStore, w, r)
	if err != nil {
		log.Println(err)
		http.Error(w, "filters error", http.StatusInternalServerError)
		return
	}

	var outfit *model.Outfit
	if f.outfitID != "" {
		outfit, err = db.SelectOutfitByID(databaseConnection, f.outfitID, userID)
		if err != nil {
			log.Println(err)
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}
	}

	itemCategories, err := service.GroupItemsByCategory(databaseConnection, userID, f.category, f.brand, f.color)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	categories, err := db.SelectCategoriesByUserID(databaseConnection, userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	brands, err := db.SelectBrandsByUserID(databaseConnection, userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	colors, err := db.SelectColorsByUserID(databaseConnection, userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	page := ui.ItemsPage{
		Page: commonui.Page{
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
	if err := template.Render(w, "items", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}

func parseFilter(sessionStore *session.Store, r *http.Request, name string) (string, error) {
	filterFromSession, err := sessionStore.Value(r, name)
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

func handleParam(sessionStore *session.Store, w http.ResponseWriter, r *http.Request, name string) (string, error) {
	value, err := parseFilter(sessionStore, r, name)
	if err != nil {
		return "", errors.Annotatef(err, "reading %s from cookie failed", name)
	}

	if sessionErr := sessionStore.SetValue(w, r, name, value); sessionErr != nil {
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

func newFilters(sessionStore *session.Store, w http.ResponseWriter, r *http.Request) (*filters, error) {
	var result filters
	var err error

	result.category, err = handleParam(sessionStore, w, r, "category")
	if err != nil {
		return nil, errors.Annotate(err, "handling category param failed")
	}

	result.brand, err = handleParam(sessionStore, w, r, "brand")
	if err != nil {
		return nil, errors.Annotate(err, "handling brand param failed")
	}

	result.color, err = handleParam(sessionStore, w, r, "color")
	if err != nil {
		return nil, errors.Annotate(err, "handling color param failed")
	}

	result.outfitID, err = handleParam(sessionStore, w, r, addToOutfitID)
	if err != nil {
		return nil, errors.Annotate(err, "handling outfit ID param failed")
	}

	return &result, nil
}
