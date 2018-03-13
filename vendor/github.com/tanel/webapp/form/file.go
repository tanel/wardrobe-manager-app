package form

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/juju/errors"
)

// File parses a file from form
func File(r *http.Request, name string) ([]byte, error) {
	file, _, err := r.FormFile(name)
	if err != nil && err != http.ErrMissingFile {
		return nil, errors.Annotate(err, "getting form file failed")
	}

	if file == nil {
		return nil, nil
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Println(errors.Annotate(closeErr, "closing file failed"))
		}
	}()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.Annotate(err, "reading form file failed")
	}

	return b, nil
}
