package db

import (
	"database/sql"

	"github.com/juju/errors"
	"github.com/tanel/wardrobe-manager-app/model"
)

func InsertItem(item model.Item) error {
	_, err := db.Exec(`
		INSERT INTO items(
			id,
			user_id,
			name,
			description,
			color,
			size,
			brand,
			price,
			currency,
			category,
			season,
			formal,
			created_at
		) VALUES(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11,
			$12,
			$13
		)
	`,
		item.ID,
		item.UserID,
		item.Name,
		item.Description,
		item.Color,
		item.Size,
		item.Brand,
		item.Price,
		item.Currency,
		item.Category,
		item.Season,
		item.Formal,
		item.CreatedAt,
	)
	if err != nil {
		return errors.Annotate(err, "inserting item failed")
	}

	return nil
}

func SelectItemWithImagesByID(itemID, userID string) (*model.Item, error) {
	item, err := SelectItemByID(itemID, userID)
	if err != nil {
		return nil, errors.Annotate(err, "selecting item failed")
	}

	itemImages, err := SelectItemImagesByItemID(itemID)
	if err != nil {
		return nil, errors.Annotate(err, "selecting item images failed")
	}

	item.Images = itemImages

	return item, nil
}

func SelectItemByID(itemID, userID string) (*model.Item, error) {
	var description, color, size, brand, currency, category sql.NullString
	var price sql.NullFloat64

	var item model.Item
	err := db.QueryRow(`
		SELECT
			id,
			user_id,
			name,
			description,
			color,
			size,
			brand,
			price,
			currency,
			category,
			season,
			formal,
			created_at
		FROM
			items
		WHERE
			id = $1
		AND
			user_id = $2
	`,
		itemID,
		userID,
	).Scan(
		&item.ID,
		&item.UserID,
		&item.Name,
		&description,
		&color,
		&size,
		&brand,
		&price,
		&currency,
		&category,
		&item.Season,
		&item.Formal,
		&item.CreatedAt,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Annotate(err, "selecting item by ID failed")
	}

	item.Description = description.String
	item.Color = color.String
	item.Size = size.String
	item.Brand = brand.String
	item.Price = price.Float64
	item.Currency = currency.String
	item.Category = category.String

	return &item, nil
}

func SelectItemsByUserID(userID string) ([]model.Item, error) {
	rows, err := db.Query(`
		SELECT
			id,
			user_id,
			name,
			description,
			color,
			size,
			brand,
			price,
			currency,
			category,
			season,
			formal,
			created_at,
			(
				SELECT
					id
				FROM
					item_images
				WHERE
					item_id = items.id
				ORDER BY
					created_at
				LIMIT 1
			) AS image_id
		FROM
			items
		WHERE
			user_id = $1
		AND
			deleted_at IS NULL
		ORDER BY
			COALESCE(category, ''), created_at, name
	`,
		userID,
	)
	if err != nil {
		return nil, errors.Annotate(err, "selecting items by user ID failed")
	}

	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		var description, color, size, brand, currency, category sql.NullString
		var price sql.NullFloat64

		var item model.Item
		if err := rows.Scan(
			&item.ID,
			&item.UserID,
			&item.Name,
			&description,
			&color,
			&size,
			&brand,
			&price,
			&currency,
			&category,
			&item.Season,
			&item.Formal,
			&item.CreatedAt,
			&item.ImageID,
		); err != nil {
			return nil, errors.Annotate(err, "scanning items failed")
		}

		item.Description = description.String
		item.Color = color.String
		item.Size = size.String
		item.Brand = brand.String
		item.Price = price.Float64
		item.Currency = currency.String
		item.Category = category.String

		items = append(items, item)
	}

	return items, nil
}

func UpdateItem(item model.Item) error {
	if item.ID == "" {
		return errors.New("item is missing ID")
	}

	_, err := db.Exec(`
		UPDATE
			items
		SET
			name = $1,
			description = $2,
			color = $3,
			size = $4,
			brand = $5,
			price = $6,
			currency = $7,
			category = $8,
			season = $9,
			formal = $10
		WHERE
			id = $11
	`,
		item.Name,
		item.Description,
		item.Color,
		item.Size,
		item.Brand,
		item.Price,
		item.Currency,
		item.Category,
		item.Season,
		item.Formal,
		item.ID,
	)
	if err != nil {
		return errors.Annotate(err, "updating item failed")
	}

	return nil
}

func DeleteItem(itemID, userID string) error {
	_, err := db.Exec(`
		UPDATE
			items
		SET
			deleted_at = current_timestamp
		WHERE
			id = $1
		AND
			user_id = $2
	`,
		itemID,
		userID,
	)
	if err != nil {
		return errors.Annotate(err, "deleting item failed")
	}

	return nil
}
