package form

import (
	"net/http"
	"strconv"

	"github.com/juju/errors"
)

// Bool parses a bool from form
func Bool(r *http.Request, name string) (bool, error) {
	boolValue, err := strconv.ParseBool(r.FormValue(name))
	if err != nil {
		return false, errors.Annotate(err, "parsing bool failed")
	}

	return boolValue, nil
}
