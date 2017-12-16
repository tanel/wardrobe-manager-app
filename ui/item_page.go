package ui

import (
	"github.com/tanel/wardrobe-manager-app/model"
)

type ItemPage struct {
	Page
	Item       model.Item
	Colors     []string
	Categories []string
}

func NewItemPage(userID string, item model.Item) *ItemPage {
	page := ItemPage{
		Page:       *NewPage(userID),
		Item:       item,
		Colors:     colors,
		Categories: categories,
	}

	return &page
}
