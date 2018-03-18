package model

import (
	"github.com/tanel/webapp/model"
)

// ItemImage represents an image of an item
type ItemImage struct {
	model.Image
	ItemID string
}
