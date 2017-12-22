package db

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-manager-app/model"
)

// InsertWeight inserts a weight into database
func InsertWeight(weightEntry model.WeightEntry) error {
	_, err := db.Exec(`
		INSERT INTO weight_entries(
			id,
			user_id,
			value,
			created_at
		) VALUES(
			$1,
			$2,
			$3,
		)
	`,
		weightEntry.ID,
		weightEntry.UserID,
		weightEntry.Value,
		weightEntry.CreatedAt,
	)
	if err != nil {
		return errors.Annotate(err, "inserting weight failed")
	}

	return nil
}
