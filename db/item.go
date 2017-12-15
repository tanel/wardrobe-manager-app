package db

import (
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
			formal
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
			$12
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
	)
	if err != nil {
		return errors.Annotate(err, "inserting item failed")
	}

	return nil
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
			formal
		FROM
			items
		WHERE
			user_id = $1
		ORDER BY
			name
	`,
		userID,
	)
	if err != nil {
		return nil, errors.Annotate(err, "selecting items by user ID failed")
	}

	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		var item model.Item
		if err := rows.Scan(
			&item.ID,
			&item.UserID,
			&item.Name,
			&item.Description,
			&item.Color,
			&item.Size,
			&item.Brand,
			&item.Price,
			&item.Currency,
			&item.Category,
			&item.Season,
			&item.Formal,
		); err != nil {
			return nil, errors.Annotate(err, "scanning items failed")
		}

		items = append(items, item)
	}

	return items, nil
}
