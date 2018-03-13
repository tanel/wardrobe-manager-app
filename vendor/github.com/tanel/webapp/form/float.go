package form

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/juju/errors"
)

// Float parses a float from form
func Float(r *http.Request, name string) (float64, error) {
	s := strings.TrimSpace(r.FormValue(name))
	if s == "" {
		return 0, nil
	}

	s = strings.Replace(s, ",", ".", -1)
	floatValue, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, errors.Annotate(err, "parsing float failed")
	}

	return floatValue, nil
}
