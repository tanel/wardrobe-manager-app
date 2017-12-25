package ui

import (
	"github.com/tanel/wardrobe-manager-app/model"
)

// OutfitsPage represents an outfits page
type OutfitsPage struct {
	Page
	Outfits []model.Outfit
}
