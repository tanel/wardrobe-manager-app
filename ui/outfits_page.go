package ui

import (
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/webapp/ui"
)

// OutfitsPage represents an outfits page
type OutfitsPage struct {
	ui.Page
	Outfits []model.Outfit
}
