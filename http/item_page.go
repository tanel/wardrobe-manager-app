package http

import (
	"github.com/tanel/wardrobe-manager-app/model"
)

type ItemPage struct {
	Page
	Item model.Item
}
