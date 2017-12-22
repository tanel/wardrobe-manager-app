package model

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/juju/errors"
)

// ItemImage represents an image of an item
type ItemImage struct {
	Base
	ItemID string
	UserID string
	Body   []byte
}

// DirectoryPath returns upload path
func (itemImage ItemImage) DirectoryPath() string {
	return filepath.Join("uploads", itemImage.UserID, "item-images")
}

// FilePath returns the full path of the image file
func (itemImage ItemImage) FilePath() string {
	directoryPath := itemImage.DirectoryPath()
	return filepath.Join(directoryPath, itemImage.ID)
}

// Save saves image to disk
func (itemImage ItemImage) Save() error {
	directoryPath := itemImage.DirectoryPath()
	if err := os.MkdirAll(directoryPath, 0700); err != nil && !strings.Contains(err.Error(), "file exists") {
		return errors.Annotate(err, "creating image directory failed")
	}

	filePath := itemImage.FilePath()
	if err := ioutil.WriteFile(filePath, itemImage.Body, 0644); err != nil {
		return errors.Annotate(err, "writing image failed")
	}

	return nil
}

// Load loads image from disk
func (itemImage *ItemImage) Load() error {
	filePath := itemImage.FilePath()
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.Annotate(err, "writing image failed")
	}

	itemImage.Body = b

	return nil
}
