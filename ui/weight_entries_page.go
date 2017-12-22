package ui

import (
	"github.com/tanel/wardrobe-manager-app/model"
)

// WeightEntriesPage represents weights page
type WeightEntriesPage struct {
	Page
	Weights []model.WeightEntry
}

// NewWeightEntriesPage returns a new weights page
func NewWeightEntriesPage(userID string, weights []model.WeightEntry) *WeightEntriesPage {
	page := WeightEntriesPage{
		Page:    *NewPage(userID),
		Weights: weights,
	}

	return &page
}
