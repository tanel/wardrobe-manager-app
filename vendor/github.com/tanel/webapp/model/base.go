package model

import (
	"time"
)

// Base represents a base model with common fields
type Base struct {
	ID        string
	CreatedAt time.Time
	DeletedAt *time.Time
}
