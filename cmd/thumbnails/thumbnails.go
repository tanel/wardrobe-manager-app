package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/juju/errors"
	"gopkg.in/h2non/bimg.v1"
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
		if err := generateThumbnailsForImage(path, 140, 200); err != nil {
			return errors.Annotatef(err, "generating thumbnail for image %s failed", path)
		}
	}

	return nil
}

func generateThumbnailsForImage(imagePath string, height uint, width uint) error {
	if strings.Contains(imagePath, "-thumbnail") {
		return nil
	}

	log.Println("generating thumbnail for", imagePath)

	buffer, err := bimg.Read(imagePath)
	if err != nil {
		return errors.Annotate(err, "reading image failed")
	}

	newImage, err := bimg.NewImage(buffer).Resize(140, 200)
	if err != nil {
		return errors.Annotate(err, "creating new image failed")
	}

	bimg.Write(imagePath+"-thumbnail", newImage)

	return nil
}
