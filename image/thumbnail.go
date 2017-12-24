package image

import (
	"log"
	"strings"

	"github.com/juju/errors"
	"gopkg.in/h2non/bimg.v1"
)

func GenerateThumbnailsForImage(imagePath string) error {
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
