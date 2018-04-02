package model

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/juju/errors"
)

// Image represents an image
type Image struct {
	Base
	UserID string `json:"user_id"`
	Body   []byte `json:"body"`
}

// DirectoryPath returns upload path
func (img Image) DirectoryPath() string {
	return filepath.Join("uploads", img.UserID, "images")
}

// FilePath returns the full path of the image file
func (img Image) FilePath() string {
	directoryPath := img.DirectoryPath()
	return filepath.Join(directoryPath, img.ID)
}

// Save saves image to disk
func (img Image) Save() error {
	directoryPath := img.DirectoryPath()
	if err := os.MkdirAll(directoryPath, 0700); err != nil && !strings.Contains(err.Error(), "file exists") {
		return errors.Annotate(err, "creating image directory failed")
	}

	filePath := img.FilePath()
	if err := ioutil.WriteFile(filePath, img.Body, 0644); err != nil {
		return errors.Annotate(err, "writing image failed")
	}

	return nil
}

// Load loads image from disk
func (img *Image) Load() error {
	filePath := img.FilePath()
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.Annotate(err, "writing image failed")
	}

	img.Body = b

	return nil
}

// LoadThumbnail loads image from disk
func (img *Image) LoadThumbnail() error {
	filePath := img.FilePath() + "-thumbnail"
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.Annotate(err, "writing image failed")
	}

	img.Body = b

	return nil
}
