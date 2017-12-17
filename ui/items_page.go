package ui

import (
	"github.com/tanel/wardrobe-manager-app/model"
)

type ItemsPage struct {
	Page
	Items      []model.Item
	Categories []model.Category
}
