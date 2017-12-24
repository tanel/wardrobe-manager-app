package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/juju/errors"
	"github.com/nfnt/resize"
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

	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal(err)
	}

	img, format, err := image.Decode(file)
	if err != nil {
		return errors.Annotate(err, "decoding image failed")
	}
	file.Close()

	log.Println(format)

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(width, 0, img, resize.Lanczos3)

	out, err := os.Create(imagePath + "-thumbnail")
	if err != nil {
		return errors.Annotate(err, "creating thumbnail image failed")
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)

	return nil
}
