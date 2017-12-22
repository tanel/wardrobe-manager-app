package service

import (
	"time"

	"github.com/juju/errors"
	"github.com/satori/go.uuid"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/model"
)

// SaveItem saves item to database, including images
func SaveItem(item *model.Item, userID string) error {
	if item.ID == "" {
		item.ID = uuid.NewV4().String()
		if err := db.InsertItem(*item); err != nil {
			return errors.Annotate(err, "inserting item failed")
		}
	} else {
		if err := db.UpdateItem(*item); err != nil {
			return errors.Annotate(err, "updating item failed")
		}
	}

	for _, itemImage := range item.Images {
		itemImage.ID = uuid.NewV4().String()
		itemImage.ItemID = item.ID
		itemImage.CreatedAt = time.Now()
		itemImage.UserID = userID
		if err := db.InsertItemImage(itemImage); err != nil {
			return errors.Annotate(err, "saving image failed")
		}

		if err := itemImage.Save(); err != nil {
			return errors.Annotate(err, "saving image file failed")
		}
	}

	return nil
}

// GroupItemsByCategory groups items into categories
func GroupItemsByCategory(userID string, category string) ([]model.Category, error) {
	items, err := db.SelectItemsByUserID(userID, category)
	if err != nil {
		return nil, errors.Annotate(err, "selectin items by user ID failed")
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
