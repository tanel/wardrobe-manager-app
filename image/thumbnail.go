package image

import (
	"log"
	"os/exec"
	"strings"

	"github.com/juju/errors"
)

// GenerateThumbnailsForImage generates a thumbnail for a given image
func GenerateThumbnailsForImage(imagePath string) error {
	if strings.Contains(imagePath, "-thumbnail") {
		return nil
	}

	log.Println("generating thumbnail for", imagePath)

	if _, err := exec.Command("cp", imagePath, imagePath+".png").Output(); err != nil {
		return errors.Annotate(err, "copying image failed")
	}

	if _, err := exec.Command("mogrify", "-format", "png", "-thumbnail", "140x140", imagePath+".png").Output(); err != nil {
		return errors.Annotate(err, "creating thumbnail failed")
	}

	if _, err := exec.Command("mv", imagePath+".png", imagePath+"-thumbnail").Output(); err != nil {
		return errors.Annotate(err, "moving thumbnail failed")
	}

	return nil
}
