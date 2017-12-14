package db

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-manager-app/model"
)

func InsertItem(item model.Item) error {
	_, err := db.Exec(`
		INSERT INTO items(
			id,
			name,
			description,
			color,
			size,
			brand,
			price,
			currency,
			category,
			season,
			format
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
			$11
		)
	`,
		item.ID,
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
