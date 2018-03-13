package form

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/juju/errors"
)

// Int parses an int from form
func Int(r *http.Request, name string) (int, error) {
	s := strings.TrimSpace(r.FormValue(name))
	if s == "" {
		return 0, nil
	}

	intValue, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.Annotate(err, "parsing integer failed")
	}

	return intValue, nil
}
