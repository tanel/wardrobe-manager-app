package ui

import (
	"github.com/tanel/wardrobe-organizer/model"
)

// ItemImagePage represents an image page
type ItemImagePage struct {
	Page
	ItemImage model.ItemImage
}

// NewItemImagePage returns new image page
func NewItemImagePage(userID string, itemImage model.ItemImage) *ItemImagePage {
	page := ItemImagePage{
		Page:      *NewPage(userID),
		ItemImage: itemImage,
	}

	return &page
}
