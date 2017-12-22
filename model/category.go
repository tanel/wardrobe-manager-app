package model

// Category represents an item category, for example: Trousers
type Category struct {
	Description string
	Items       []Item
	ItemCount   int
}
