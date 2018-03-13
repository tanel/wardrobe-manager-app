package model

import (
	"net/http"
	"time"

	"github.com/juju/errors"
	"github.com/tanel/webapp/form"
	"github.com/tanel/webapp/model"
)

// WeightEntry represents a weight measurement
type WeightEntry struct {
	model.Base
	UserID string
	Value  float64
}

// NewWeightEntryForm returns an item with values parsed from HTTP form
func NewWeightEntryForm(r *http.Request) (*WeightEntry, error) {
	if err := r.ParseForm(); err != nil {
		return nil, errors.Annotate(err, "parsing form failed")
	}

	weight, err := form.Float(r, "weight")
	if err != nil {
		return nil, errors.New("please enter a valid weight")
	}

	if int(weight) <= 0 {
		return nil, errors.New("weight cannot be zero")
	}

	var weightEntry WeightEntry
	weightEntry.Value = weight
	weightEntry.CreatedAt = time.Now()

	return &weightEntry, nil
}
