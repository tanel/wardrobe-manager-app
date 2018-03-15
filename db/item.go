package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/juju/errors"
	"github.com/satori/go.uuid"
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/webapp/image"
)

// InsertItem inserts an item into database
func InsertItem(db *sql.DB, item model.Item) error {
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
			code,
			url,
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
			$15,
			$16,
			$17
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
		item.Code,
		item.URL,
		item.CreatedAt,
	)
	if err != nil {
		return errors.Annotate(err, "inserting item failed")
	}

	return nil
}

// SelectItemWithImagesByID selects item by ID, including its images
func SelectItemWithImagesByID(db *sql.DB, itemID, userID string) (*model.Item, error) {
	item, err := SelectItemByID(db, itemID, userID)
	if err != nil {
		return nil, errors.Annotate(err, "selecting item failed")
	}

	itemImages, err := SelectItemImagesByItemID(db, itemID)
	if err != nil {
		return nil, errors.Annotate(err, "selecting item images failed")
	}

	item.Images = itemImages

	return item, nil
}

// SelectItemByID selects an item by ID
func SelectItemByID(db *sql.DB, itemID, userID string) (*model.Item, error) {
	var description, color, size, brand, currency, category, code, url sql.NullString
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
			code,
			url,
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
		&code,
		&url,
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
	item.Code = code.String
	item.URL = url.String

	return &item, nil
}

// SelectItemsByUserID selects items by user ID and category, brand, color
func SelectItemsByUserID(db *sql.DB, userID string, category, brand, color string) ([]model.Item, error) {
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
			starred,
			code,
			url
		FROM
			items
		WHERE
			user_id = $1
		AND
			deleted_at IS NULL
		AND
			(category = $2 OR $2 = '')
		AND
			(($3 <> '' AND brand IS NOT NULL AND brand = $3) OR ($3 = ''))
		AND
			(($4 <> '' AND color IS NOT NULL AND color = $4) OR ($4 = ''))
		ORDER BY
			COALESCE(category, ''), created_at, name
	`,
		userID,
		category,
		brand,
		color,
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
		var description, color, size, brand, currency, category, code, url sql.NullString
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
			&code,
			&url,
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
		item.Code = code.String
		item.URL = url.String

		items = append(items, item)
	}

	return items, nil
}

// UpdateItem updates item in database
func UpdateItem(db *sql.DB, item model.Item) error {
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
			starred = $12,
			code = $13,
			url = $14
		WHERE
			id = $15
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
		item.Code,
		item.URL,
		item.ID,
	)
	if err != nil {
		return errors.Annotate(err, "updating item failed")
	}

	return nil
}

// DeleteItem deletes an item
func DeleteItem(db *sql.DB, itemID, userID string) error {
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
func SelectBrandsByUserID(db *sql.DB, userID string) ([]string, error) {
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
func SelectColorsByUserID(db *sql.DB, userID string) ([]string, error) {
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

// SelectCategoriesByUserID selects categories by user ID
func SelectCategoriesByUserID(db *sql.DB, userID string) ([]string, error) {
	rows, err := db.Query(`
		SELECT
			DISTINCT category
		FROM
			items
		WHERE
			user_id = $1
		AND
			deleted_at IS NULL
		AND
			category IS NOT NULL
		AND
			category <> ''
		ORDER BY
			category
	`,
		userID,
	)
	if err != nil {
		return nil, errors.Annotate(err, "selecting categories by user ID failed")
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(errors.Annotate(err, "closing rows failed"))
		}
	}()

	var result []string
	for rows.Next() {
		var category string
		if err := rows.Scan(
			&category,
		); err != nil {
			return nil, errors.Annotate(err, "scanning categories failed")
		}

		result = append(result, category)
	}

	return result, nil
}

// SaveItem saves item to database, including images
func SaveItem(db *sql.DB, item *model.Item, userID string) error {
	if item.ID == "" {
		item.ID = uuid.Must(uuid.NewV4()).String()
		if err := InsertItem(db, *item); err != nil {
			return errors.Annotate(err, "inserting item failed")
		}
	} else {
		if err := UpdateItem(db, *item); err != nil {
			return errors.Annotate(err, "updating item failed")
		}
	}

	for _, itemImage := range item.Images {
		if len(itemImage.Body) == 0 {
			continue
		}

		itemImage.ID = uuid.Must(uuid.NewV4()).String()
		itemImage.ItemID = item.ID
		itemImage.CreatedAt = time.Now()
		itemImage.UserID = userID
		if err := InsertItemImage(db, itemImage); err != nil {
			return errors.Annotate(err, "saving image failed")
		}

		if err := itemImage.Save(); err != nil {
			return errors.Annotate(err, "saving image file failed")
		}

		if err := image.GenerateThumbnailsForImage(itemImage.FilePath()); err != nil {
			return errors.Annotatef(err, "generating thumbnail for image %s failed", itemImage.FilePath())
		}
	}

	return nil
}

// GroupItemsByCategory groups items into categories
func GroupItemsByCategory(db *sql.DB, userID string, category, brand, color string) ([]model.Category, error) {
	items, err := SelectItemsByUserID(db, userID, category, brand, color)
	if err != nil {
		return nil, errors.Annotate(err, "selecting items by user ID failed")
	}

	var descriptions []string
	categoryLookup := make(map[string]*model.Category)
	for _, item := range items {
		description := "Uncategorized"
		if item.Category != "" {
			description = item.Category
		}

		category, exists := categoryLookup[description]
		if !exists {
			category = &model.Category{
				Description: description,
			}

			descriptions = append(descriptions, description)
		}

		category.Items = append(category.Items, item)
		category.ItemCount = category.ItemCount + 1

		categoryLookup[description] = category
	}

	var categories []model.Category
	for _, description := range descriptions {
		category := categoryLookup[description]
		categories = append(categories, *category)
	}

	return categories, nil
}
