package ui

import (
	"github.com/tanel/wardrobe-manager-app/model"
)

type ItemsPage struct {
	Page
	Categories       []model.Category
	ItemCount        int64
	SelectedCategory string
}
