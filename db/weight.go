package db

import (
	"log"

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
			$4
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

// SelectWeightsByUserID selects weights by user ID
func SelectWeightsByUserID(userID string) ([]model.WeightEntry, error) {
	rows, err := db.Query(`
		SELECT
			id,
			user_id,
			value,
			created_at
		FROM
			weight_entries
		WHERE
			user_id = $1
		AND
			deleted_at IS NULL
		ORDER BY
			created_at
	`,
		userID,
	)
	if err != nil {
		return nil, errors.Annotate(err, "selecting weight entries by user ID failed")
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(errors.Annotate(err, "closing rows failed"))
		}
	}()

	var result []model.WeightEntry
	for rows.Next() {
		var weightEntry model.WeightEntry
		if err := rows.Scan(
			&weightEntry.ID,
			&weightEntry.UserID,
			&weightEntry.Value,
			&weightEntry.CreatedAt,
		); err != nil {
			return nil, errors.Annotate(err, "scanning weight entries failed")
		}

		result = append(result, weightEntry)
	}

	return result, nil
}
