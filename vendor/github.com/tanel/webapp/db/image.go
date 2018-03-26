package db

import (
	"github.com/juju/errors"
	"github.com/tanel/webapp/model"
)

// InsertImage inserts image into database
func InsertImage(img model.Image) error {
	_, err := SharedInstance.Exec(`
		INSERT INTO images(
			id,
			user_id,
			created_at
		) VALUES(
			$1,
			$2,
			$3
		)
	`,
		img.ID,
		img.UserID,
		img.CreatedAt,
	)
	if err != nil {
		return errors.Annotate(err, "inserting image failed")
	}

	return nil
}
