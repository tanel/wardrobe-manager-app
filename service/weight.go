package service

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/model"
)

func SaveWeight(weight float64, userID string) error {
	data := model.WeightEntry{
		UserID: userID,
		Value:  weight,
	}

	if err := db.InsertWeight(data); err != nil {
		return errors.Annotate(err, "inserting weight failed")
	}

	return nil
}
