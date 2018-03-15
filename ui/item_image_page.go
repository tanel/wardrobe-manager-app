package ui

import (
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/webapp/ui"
)

// ItemImagePage represents an image page
type ItemImagePage struct {
	ui.Page
	ItemImage model.ItemImage
}

// NewItemImagePage returns new image page
func NewItemImagePage(userID string, itemImage model.ItemImage) *ItemImagePage {
	return &ItemImagePage{
		Page:      *ui.NewPageWithUserID(userID),
		ItemImage: itemImage,
	}
}
