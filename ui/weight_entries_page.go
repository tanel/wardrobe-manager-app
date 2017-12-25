package ui

import (
	"encoding/json"

	"github.com/juju/errors"
	"github.com/tanel/wardrobe-manager-app/model"
)

// WeightEntriesPage represents weights page
type WeightEntriesPage struct {
	Page
	Weights             []model.WeightEntry
	WeightChartDataJSON string
}

// NewWeightEntriesPage returns a new weights page
func NewWeightEntriesPage(userID string, weights []model.WeightEntry) (*WeightEntriesPage, error) {
	page := WeightEntriesPage{
		Page:    *NewPage(userID),
		Weights: weights,
	}

	if err := page.prepareChartData(); err != nil {
		return nil, errors.Annotate(err, "preparing chart data failed")
	}

	return &page, nil
}

func (page *WeightEntriesPage) prepareChartData() error {
	var data []float64
	for _, weightEntry := range page.Weights {
		data = append(data, weightEntry.Value)
	}

	b, err := json.Marshal(data)
	if err != nil {
		return errors.Annotate(err, "marshalling data to JSON failed")
	}

	page.WeightChartDataJSON = string(b)

	return nil
}
