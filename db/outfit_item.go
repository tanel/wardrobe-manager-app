package db

import (
	"database/sql"
	"log"

	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/webapp/db"
)

// SelectOutfitIDByOutfitItemID selects outfit ID by outfit item ID
func SelectOutfitIDByOutfitItemID(outfitItemID, userID string) (string, error) {
	var outfitID string
	err := db.QueryRow(`
		SELECT
			outfit_id
		FROM
			outfit_items
		WHERE
			id = $1
		AND
			outfit_id IN (
				SELECT
					id
				FROM
					outfits
				WHERE
					user_id = $2
				AND
					deleted_at IS NULL
			)
		AND
			deleted_at IS NULL
		LIMIT 1
	`,
		outfitItemID,
		userID,
	).Scan(
		&outfitID,
	)
	if err != nil && err != sql.ErrNoRows {
		return "", errors.Annotate(err, "selecting outfit ID by outfit item ID failed")
	}

	return outfitID, nil
}

// SelectOutfitItemsByOutfitID selects outfit items by outfit ID
func SelectOutfitItemsByOutfitID(outfitID string, userID string) ([]model.OutfitItem, error) {
	rows, err := db.Query(`
		SELECT
			outfit_items.id AS outfit_item_id,
			items.id,
			items.user_id,
			items.name,
			items.description,
			items.color,
			items.size,
			items.brand,
			items.price,
			items.currency,
			items.category,
			items.season,
			items.formal,
			items.created_at,
			(
				SELECT
					id
				FROM
					item_images
				WHERE
					item_id = items.id
				AND
					deleted_at IS NULL
				ORDER BY
					created_at
				LIMIT 1
			) AS image_id,
			items.quantity,
			items.starred
		FROM
			outfit_items
		LEFT OUTER JOIN items ON items.id = outfit_items.item_id
		WHERE
			items.user_id = $1
		AND
			outfit_items.outfit_id = $2
		AND
			outfit_items.deleted_at IS NULL
		AND
			items.deleted_at IS NULL
		ORDER BY
			COALESCE(items.category, ''), outfit_items.created_at, items.name
	`,
		userID,
		outfitID,
	)
	if err != nil {
		return nil, errors.Annotate(err, "selecting outfit items by outfit ID failed")
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(errors.Annotate(err, "closing rows failed"))
		}
	}()

	var items []model.OutfitItem
	for rows.Next() {
		var description, color, size, brand, currency, category sql.NullString
		var price sql.NullFloat64

		var outfitItem model.OutfitItem
		outfitItem.Item = &model.Item{}
		if err := rows.Scan(
			&outfitItem.ID,
			&outfitItem.Item.ID,
			&outfitItem.Item.UserID,
			&outfitItem.Item.Name,
			&description,
			&color,
			&size,
			&brand,
			&price,
			&currency,
			&category,
			&outfitItem.Item.Season,
			&outfitItem.Item.Formal,
			&outfitItem.Item.CreatedAt,
			&outfitItem.Item.ImageID,
			&outfitItem.Item.Quantity,
			&outfitItem.Item.Starred,
		); err != nil {
			return nil, errors.Annotate(err, "scanning outfit items failed")
		}

		outfitItem.Item.Description = description.String
		outfitItem.Item.Color = color.String
		outfitItem.Item.Size = size.String
		outfitItem.Item.Brand = brand.String
		outfitItem.Item.Price = price.Float64
		outfitItem.Item.Currency = currency.String
		outfitItem.Item.Category = category.String

		items = append(items, outfitItem)
	}

	return items, nil
}

// InsertOutfitItem inserts an outfit item into database
func InsertOutfitItem(item model.OutfitItem) error {
	_, err := db.Exec(`
		INSERT INTO outfit_items(
			id,
			outfit_id,
			item_id,
			created_at
		) VALUES(
			$1,
			$2,
			$3,
			$4
		)
	`,
		item.ID,
		item.OutfitID,
		item.ItemID,
		item.CreatedAt,
	)
	if err != nil {
		return errors.Annotate(err, "inserting outfit item failed")
	}

	return nil
}

// DeleteOutfitItem deletes an outfit item
func DeleteOutfitItem(outfitItemID, userID string) error {
	_, err := db.Exec(`
		UPDATE
			outfit_items
		SET
			deleted_at = current_timestamp
		WHERE
			id = $1
		AND
			outfit_id IN (
				SELECT
					id
				FROM
					outfits
				WHERE
					user_id = $2
				AND
					deleted_at IS NULL
			)
	`,
		outfitItemID,
		userID,
	)
	if err != nil {
		return errors.Annotate(err, "deleting outfit item failed")
	}

	return nil
}
