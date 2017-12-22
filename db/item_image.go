package db

import (
	"log"

	"github.com/juju/errors"
	"github.com/tanel/wardrobe-manager-app/model"
)

// InsertItemImage inserts image into database
func InsertItemImage(itemImage model.ItemImage) error {
	_, err := db.Exec(`
		INSERT INTO item_images(
			id,
			item_id,
			user_id,
			created_at
		) VALUES(
			$1,
			$2,
			$3,
			$4
		)
	`,
		itemImage.ID,
		itemImage.ItemID,
		itemImage.UserID,
		itemImage.CreatedAt,
	)
	if err != nil {
		return errors.Annotate(err, "inserting item image failed")
	}

	return nil
}

// SelectItemImagesByItemID selects images by item ID
func SelectItemImagesByItemID(itemID string) ([]model.ItemImage, error) {
	rows, err := db.Query(`
		SELECT
			id,
			item_id,
			user_id,
			created_at
		FROM
			item_images
		WHERE
			item_id = $1
		AND
			deleted_at is null
		ORDER BY
			created_at
	`,
		itemID,
	)
	if err != nil {
		return nil, errors.Annotate(err, "selecting item images by item ID failed")
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(errors.Annotate(err, "closing rows failed"))
		}
	}()

	var itemImages []model.ItemImage
	for rows.Next() {
		var itemImage model.ItemImage
		if err := rows.Scan(
			&itemImage.ID,
			&itemImage.ItemID,
			&itemImage.UserID,
			&itemImage.CreatedAt,
		); err != nil {
			return nil, errors.Annotate(err, "scanning item images failed")
		}

		itemImages = append(itemImages, itemImage)
	}

	return itemImages, nil
}

// SelectItemImageByID selects an image by ID
func SelectItemImageByID(itemImageID, userID string) (*model.ItemImage, error) {
	var itemImage model.ItemImage
	if err := db.QueryRow(`
		SELECT
			id,
			item_id,
			user_id,
			created_at
		FROM
			item_images
		WHERE
			id = $1
		AND
			user_id = $2
		LIMIT 1
	`,
		itemImageID,
		userID,
	).Scan(
		&itemImage.ID,
		&itemImage.ItemID,
		&itemImage.UserID,
		&itemImage.CreatedAt,
	); err != nil {
		return nil, errors.Annotate(err, "scanning item image failed")
	}

	return &itemImage, nil
}

// DeleteItemImage deletes an item in database
func DeleteItemImage(itemImageID, userID string) error {
	_, err := db.Exec(`
		UPDATE
			item_images
		SET
			deleted_at = current_timestamp
		WHERE
			id = $1
		AND
			user_id = $2
	`,
		itemImageID,
		userID,
	)
	if err != nil {
		return errors.Annotate(err, "deleting item image failed")
	}

	return nil
}
