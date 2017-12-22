package ui

import (
	"github.com/tanel/wardrobe-manager-app/model"
)

// ItemsPage represents an items page
type ItemsPage struct {
	Page
	Categories       []model.Category
	SelectedCategory string
}
