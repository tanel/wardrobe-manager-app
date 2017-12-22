package ui

import (
	"github.com/tanel/wardrobe-manager-app/model"
)

// ItemPage represents item page
type ItemPage struct {
	Page
	Item model.Item
}

// NewItemPage returns a new item page
func NewItemPage(userID string, item model.Item) *ItemPage {
	page := ItemPage{
		Page: *NewPage(userID),
		Item: item,
	}

	return &page
}
