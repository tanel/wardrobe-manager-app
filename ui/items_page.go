package ui

import (
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/webapp/ui"
)

// ItemsPage represents an items page
type ItemsPage struct {
	ui.Page
	ItemCategories   []model.Category
	Categories       []string
	Brands           []string
	Colors           []string
	SelectedCategory string
	SelectedBrand    string
	SelectedColor    string
	SelectedOutfit   *model.Outfit
}
