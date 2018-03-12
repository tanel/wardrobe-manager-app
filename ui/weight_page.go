package ui

import (
	"github.com/tanel/wardrobe-organizer/model"
)

// WeightPage represents weight page
type WeightPage struct {
	Page
	Weight model.WeightEntry
}

// NewWeightPage returns a new weight page
func NewWeightPage(userID string, weight model.WeightEntry) *WeightPage {
	page := WeightPage{
		Page:   *NewPage(userID),
		Weight: weight,
	}

	return &page
}
