package ui

import (
	"github.com/tanel/wardrobe-manager-app/model"
)

// OutfitPage represents outfit page
type OutfitPage struct {
	Page
	Outfit model.Outfit
}

// NewOutfitPage returns a new outfit page
func NewOutfitPage(userID string, outfit model.Outfit) *OutfitPage {
	page := OutfitPage{
		Page:   *NewPage(userID),
		Outfit: outfit,
	}

	return &page
}
