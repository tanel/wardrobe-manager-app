package service

import (
	"database/sql"
	"time"

	"github.com/juju/errors"
	"github.com/satori/go.uuid"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/webapp/image"
)

// SaveItem saves item to database, including images
func SaveItem(connection *sql.DB, item *model.Item, userID string) error {
	if item.ID == "" {
		item.ID = uuid.Must(uuid.NewV4()).String()
		if err := db.InsertItem(connection, *item); err != nil {
			return errors.Annotate(err, "inserting item failed")
		}
	} else {
		if err := db.UpdateItem(connection, *item); err != nil {
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
		if err := db.InsertItemImage(connection, itemImage); err != nil {
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
func GroupItemsByCategory(connection *sql.DB, userID string, category, brand, color string) ([]model.Category, error) {
	items, err := db.SelectItemsByUserID(connection, userID, category, brand, color)
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
