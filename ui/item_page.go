package ui

import (
	"github.com/tanel/wardrobe-manager-app/model"
)

type ItemPage struct {
	Page
	Item model.Item
}

func NewItemPage(userID string, item model.Item) *ItemPage {
	page := ItemPage{
		Page: *NewPage(userID),
		Item: item,
	}

	return &page
}
