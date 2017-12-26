package model

// OutfitItem represents an item that belongs to an outfit
type OutfitItem struct {
	Base
	ItemID   string
	OutfitID string

	Item *Item
}
