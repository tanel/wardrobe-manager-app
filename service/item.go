package service

import (
	"time"

	"github.com/juju/errors"
	"github.com/satori/go.uuid"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/model"
)

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
		if err := db.InsertItemImage(itemImage); err != nil {
			return errors.Annotate(err, "saving image failed")
		}

		if err := itemImage.Save(userID); err != nil {
			return errors.Annotate(err, "saving image file failed")
		}
	}

	return nil
}
