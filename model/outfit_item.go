package model

import (
	"github.com/tanel/webapp/model"
)

// OutfitItem represents an item that belongs to an outfit
type OutfitItem struct {
	model.Base
	ItemID   string
	OutfitID string

	Item *Item
}
