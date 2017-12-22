package model

// WeightEntry represents a weight measurement
type WeightEntry struct {
	Base
	UserID string
	Value  float64
}
