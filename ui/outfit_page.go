package ui

import (
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/webapp/ui"
)

// OutfitPage represents outfit page
type OutfitPage struct {
	ui.Page
	Outfit model.Outfit
}

// NewOutfitPage returns a new outfit page
func NewOutfitPage(userID string, outfit model.Outfit) *OutfitPage {
	page := OutfitPage{
		Page:   *ui.NewPage(userID),
		Outfit: outfit,
	}

	return &page
}
