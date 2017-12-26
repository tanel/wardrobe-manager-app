package db

import (
	"database/sql"
	"log"

	"github.com/juju/errors"
	"github.com/tanel/wardrobe-manager-app/model"
)

// InsertItem inserts an item into database
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
			quantity,
			starred,
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
			$13,
			$14,
			$15
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
		item.Quantity,
		item.Starred,
		item.CreatedAt,
	)
	if err != nil {
		return errors.Annotate(err, "inserting item failed")
	}

	return nil
}

// SelectItemWithImagesByID selects item by ID, including its images
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

// SelectItemByID selects an item by ID
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
			quantity,
			starred,
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
		&item.Quantity,
		&item.Starred,
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

// SelectItemsByUserID selects items by user ID and category
func SelectItemsByUserID(userID string, category string) ([]model.Item, error) {
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
				AND
					deleted_at IS NULL
				ORDER BY
					created_at
				LIMIT 1
			) AS image_id,
			quantity,
			starred
		FROM
			items
		WHERE
			user_id = $1
		AND
			deleted_at IS NULL
		AND
			(category = $2 OR $2 = '')
		ORDER BY
			COALESCE(category, ''), created_at, name
	`,
		userID,
		category,
	)
	if err != nil {
		return nil, errors.Annotate(err, "selecting items by user ID failed")
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(errors.Annotate(err, "closing rows failed"))
		}
	}()

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
			&item.Quantity,
			&item.Starred,
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

// UpdateItem updates item in database
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
			formal = $10,
			quantity = $11,
			starred = $12
		WHERE
			id = $13
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
		item.Quantity,
		item.Starred,
		item.ID,
	)
	if err != nil {
		return errors.Annotate(err, "updating item failed")
	}

	return nil
}

// DeleteItem deletes an item
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

// SelectBrandsByUserID selects brands by user ID
func SelectBrandsByUserID(userID string) ([]string, error) {
	rows, err := db.Query(`
		SELECT
			DISTINCT brand
		FROM
			items
		WHERE
			user_id = $1
		AND
			deleted_at IS NULL
		AND
			brand IS NOT NULL
		AND
			brand <> ''
		ORDER BY
			brand
	`,
		userID,
	)
	if err != nil {
		return nil, errors.Annotate(err, "selecting brands by user ID failed")
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(errors.Annotate(err, "closing rows failed"))
		}
	}()

	var result []string
	for rows.Next() {
		var brand string
		if err := rows.Scan(
			&brand,
		); err != nil {
			return nil, errors.Annotate(err, "scanning brands failed")
		}

		result = append(result, brand)
	}

	return result, nil
}

// SelectColorsByUserID selects colors by user ID
func SelectColorsByUserID(userID string) ([]string, error) {
	rows, err := db.Query(`
		SELECT
			DISTINCT color
		FROM
			items
		WHERE
			user_id = $1
		AND
			deleted_at IS NULL
		AND
			color IS NOT NULL
		AND
			color <> ''
		ORDER BY
			color
	`,
		userID,
	)
	if err != nil {
		return nil, errors.Annotate(err, "selecting colors by user ID failed")
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(errors.Annotate(err, "closing rows failed"))
		}
	}()

	var result []string
	for rows.Next() {
		var color string
		if err := rows.Scan(
			&color,
		); err != nil {
			return nil, errors.Annotate(err, "scanning colors failed")
		}

		result = append(result, color)
	}

	return result, nil
}
