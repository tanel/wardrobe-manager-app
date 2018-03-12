package ui

import (
	"github.com/tanel/wardrobe-organizer/model"
)

// ItemsPage represents an items page
type ItemsPage struct {
	Page
	ItemCategories   []model.Category
	Categories       []string
	Brands           []string
	Colors           []string
	SelectedCategory string
	SelectedBrand    string
	SelectedColor    string
	SelectedOutfit   *model.Outfit
}
