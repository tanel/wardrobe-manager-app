package model

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/juju/errors"
)

type ItemImage struct {
	Base
	ItemID string
	UserID string
	Body   []byte
}

func (itemImage ItemImage) DirectoryPath() string {
	return filepath.Join("uploads", itemImage.UserID, "item-images")
}

func (itemImage ItemImage) FilePath() string {
	directoryPath := itemImage.DirectoryPath()
	return filepath.Join(directoryPath, itemImage.ID)
}

func (itemImage ItemImage) SaveImages() error {
	directoryPath := itemImage.DirectoryPath()
	if err := os.MkdirAll(directoryPath, 0777); err != nil && !strings.Contains(err.Error(), "file exists") {
		return errors.Annotate(err, "creating image directory failed")
	}

	filePath := itemImage.FilePath()
	if err := ioutil.WriteFile(filePath, itemImage.Body, 0644); err != nil {
		return errors.Annotate(err, "writing image failed")
	}

	return nil
}

func (itemImage *ItemImage) LoadImages() error {
	filePath := itemImage.FilePath()
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.Annotate(err, "writing image failed")
	}

	itemImage.Body = b

	return nil
}
