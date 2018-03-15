package controller

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/wardrobe-organizer/ui"
	"github.com/tanel/webapp/http"
	commonui "github.com/tanel/webapp/ui"
)

// GetItems renders items page
func GetItems(request *http.Request, userID string) {
	f, ok := newFilters(request)
	if !ok {
		return
	}

	var outfit *model.Outfit
	if f.outfitID != "" {
		var err error
		outfit, err = db.SelectOutfitByID(request.DB, f.outfitID, userID)
		if err != nil {
			request.InternalServerError(errors.Annotate(err, "selecting outfit by ID failed"))
			return
		}
	}

	itemCategories, err := db.GroupItemsByCategory(request.DB, userID, f.category, f.brand, f.color)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "grouping items by category failed"))
		return
	}

	categories, err := db.SelectCategoriesByUserID(request.DB, userID)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "selecting categories by user ID failed"))
		return
	}

	brands, err := db.SelectBrandsByUserID(request.DB, userID)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "selecting brands by user ID failed"))
		return
	}

	colors, err := db.SelectColorsByUserID(request.DB, userID)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "selecting colors by user ID failed"))
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
	request.Render("items", page)
}

func parseFilter(request *http.Request, name string) (string, bool) {
	filterFromSession, ok := request.SessionValue(name)
	if !ok {
		return "", false
	}

	filterFromURL := request.QueryParamByName(name)
	hasFilterParameter := false
	for k := range request.Query() {
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

	return filter, true
}

func handleParam(request *http.Request, name string) (string, bool) {
	value, ok := parseFilter(request, name)
	if !ok {
		return "", false
	}

	if ok := request.SetSessionValue(name, value); !ok {
		return "", false
	}

	return value, true
}

type filters struct {
	category string
	brand    string
	color    string
	outfitID string
}

func newFilters(request *http.Request) (*filters, bool) {
	var result filters
	var ok bool

	result.category, ok = handleParam(request, "category")
	if !ok {
		return nil, false
	}

	result.brand, ok = handleParam(request, "brand")
	if !ok {
		return nil, false
	}

	result.color, ok = handleParam(request, "color")
	if !ok {
		return nil, false
	}

	result.outfitID, ok = handleParam(request, addToOutfitID)
	if !ok {
		return nil, false
	}

	return &result, true
}
