package model

import (
	"time"
)

type ItemImage struct {
	ID        string
	ItemID    string
	UserID    string
	CreatedAt time.Time
	Body      []byte
}
