package db

import (
	"database/sql"
	"log"

	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/model"
)

// InsertOutfit inserts an outfit into database
func InsertOutfit(outfit model.Outfit) error {
	_, err := db.Exec(`
		INSERT INTO outfits(
			id,
			user_id,
			name,
			created_at
		) VALUES(
			$1,
			$2,
			$3,
			$4
		)
	`,
		outfit.ID,
		outfit.UserID,
		outfit.Name,
		outfit.CreatedAt,
	)
	if err != nil {
		return errors.Annotate(err, "inserting outfit failed")
	}

	return nil
}

// SelectOutfitByID selects a outfit by ID
func SelectOutfitByID(outfitID, userID string) (*model.Outfit, error) {
	var outfit model.Outfit
	err := db.QueryRow(`
		SELECT
			id,
			user_id,
			name,
			created_at
		FROM
			outfits
		WHERE
			id = $1
		AND
			user_id = $2
	`,
		outfitID,
		userID,
	).Scan(
		&outfit.ID,
		&outfit.UserID,
		&outfit.Name,
		&outfit.CreatedAt,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Annotate(err, "selecting outfit by ID failed")
	}

	return &outfit, nil
}

// UpdateOutfit updates an outfit in database
func UpdateOutfit(outfit model.Outfit) error {
	_, err := db.Exec(`
		UPDATE
			outfits
		SET
			name = $1
		WHERE
			id = $2
		AND
			user_id = $3
	`,
		outfit.Name,
		outfit.ID,
		outfit.UserID,
	)
	if err != nil {
		return errors.Annotate(err, "updating outfit failed")
	}

	return nil
}

// SelectOutfitsByUserID selects outfits by user ID
func SelectOutfitsByUserID(userID string) ([]model.Outfit, error) {
	rows, err := db.Query(`
		SELECT
			id,
			user_id,
			name,
			created_at
		FROM
			outfits
		WHERE
			user_id = $1
		AND
			deleted_at IS NULL
		ORDER BY
			created_at DESC
	`,
		userID,
	)
	if err != nil {
		return nil, errors.Annotate(err, "selecting outfits by user ID failed")
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(errors.Annotate(err, "closing rows failed"))
		}
	}()

	var result []model.Outfit
	for rows.Next() {
		var outfit model.Outfit
		if err := rows.Scan(
			&outfit.ID,
			&outfit.UserID,
			&outfit.Name,
			&outfit.CreatedAt,
		); err != nil {
			return nil, errors.Annotate(err, "scanning outfits failed")
		}

		result = append(result, outfit)
	}

	return result, nil
}

// DeleteOutfit delets an outfit in database
func DeleteOutfit(outfitID, userID string) error {
	_, err := db.Exec(`
		UPDATE
			outfits
		SET
			deleted_at = current_timestamp
		WHERE
			id = $1
		AND
			user_id = $2
	`,
		outfitID,
		userID,
	)
	if err != nil {
		return errors.Annotate(err, "deleting outfit failed")
	}

	return nil
}
