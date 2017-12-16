package model

import (
	"time"
)

type ItemImage struct {
	ID        string
	ItemID    string
	CreatedAt time.Time
	Body      []byte
}
