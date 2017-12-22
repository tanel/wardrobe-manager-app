package model

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/juju/errors"
)

// WeightEntry represents a weight measurement
type WeightEntry struct {
	Base
	UserID string
	Value  float64
}

// WeightEntryForm returns an item with values parsed from HTTP form
func NewWeightEntryForm(r *http.Request) (*WeightEntry, error) {
	if err := r.ParseForm(); err != nil {
		return nil, errors.Annotate(err, "parsing form failed")
	}

	var weightEntry WeightEntry

	s := strings.TrimSpace(r.FormValue("weight"))
	s = strings.Replace(s, ",", ".", -1)
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, errors.New("please enter a valid weight")
	}

	weightEntry.Value = value

	weightEntry.CreatedAt = time.Now()

	return &weightEntry, nil
}
