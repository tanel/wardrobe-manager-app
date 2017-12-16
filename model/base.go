package model

import (
	"time"
)

type Base struct {
	ID        string
	CreatedAt time.Time
	DeletedAt *time.Time
}
