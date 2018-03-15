package ui

import (
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/webapp/ui"
)

// ItemPage represents item page
type ItemPage struct {
	ui.Page
	Item model.Item
}

// NewItemPage returns a new item page
func NewItemPage(userID string, item model.Item) *ItemPage {
	return &ItemPage{
		Page: *ui.NewPageWithUserID(userID),
		Item: item,
	}
}
