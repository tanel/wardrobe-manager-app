package db

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-manager-app/model"
)

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
		ORDER BY
			created_at
	`,
		itemID,
	)
	if err != nil {
		return nil, errors.Annotate(err, "selecting item images by item ID failed")
	}

	defer rows.Close()

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
