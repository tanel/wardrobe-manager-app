package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/image"
)

func main() {
	if err := generateThumbnails(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func generateThumbnails() error {
	pattern := filepath.Join("uploads", "*")
	userFolders, err := filepath.Glob(pattern)
	if err != nil {
		return errors.Annotatef(err, "globbing %s failed", pattern)
	}

	for _, path := range userFolders {
		if err := generateThumbnailsForUserFolder(path); err != nil {
			return errors.Annotatef(err, "generating thumbnails for user folder %s failed", path)
		}
	}

	return nil
}

func generateThumbnailsForUserFolder(userFolder string) error {
	pattern := filepath.Join(userFolder, "item-images", "*")
	images, err := filepath.Glob(pattern)
	if err != nil {
		return errors.Annotatef(err, "globbing %s failed", pattern)
	}

	for _, path := range images {
		if err := image.GenerateThumbnailsForImage(path); err != nil {
			return errors.Annotatef(err, "generating thumbnail for image %s failed", path)
		}
	}

	return nil
}
