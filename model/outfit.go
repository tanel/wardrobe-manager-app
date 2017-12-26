package model

import (
	"net/http"
	"strings"

	"github.com/juju/errors"
)

// Outfit represents an outfit
type Outfit struct {
	Base
	UserID      string
	Name        string
	OutfitItems []OutfitItem
}

// NewOutfitForm returns an outfit with values parsed from HTTP form
func NewOutfitForm(r *http.Request) (*Outfit, error) {
	if err := r.ParseForm(); err != nil {
		return nil, errors.Annotate(err, "parsing form failed")
	}

	name := strings.TrimSpace(r.FormValue("name"))
	if name == "" {
		return nil, errors.New("please enter a name")
	}

	var outfit Outfit
	outfit.Name = name

	return &outfit, nil
}
