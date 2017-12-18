package ui

import (
	"github.com/tanel/wardrobe-manager-app/model"
)

type ItemImagePage struct {
	Page
	ItemImage model.ItemImage
}

func NewItemImagePage(userID string, itemImage model.ItemImage) *ItemImagePage {
	page := ItemImagePage{
		Page:      *NewPage(userID),
		ItemImage: itemImage,
	}

	return &page
}
