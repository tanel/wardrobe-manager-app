package model

import (
	"time"
)

// Base represents a base model with common fields
type Base struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
