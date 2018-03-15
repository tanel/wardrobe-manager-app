package ui

import (
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/webapp/ui"
)

// WeightPage represents weight page
type WeightPage struct {
	ui.Page
	Weight model.WeightEntry
}

// NewWeightPage returns a new weight page
func NewWeightPage(userID string, weight model.WeightEntry) *WeightPage {
	return &WeightPage{
		Page:   *ui.NewPageWithUserID(userID),
		Weight: weight,
	}
}
