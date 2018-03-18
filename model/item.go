package model

import (
	"net/http"
	"strings"
	"time"

	"github.com/juju/errors"
	"github.com/tanel/webapp/form"
	"github.com/tanel/webapp/model"
)

// Item represents a wardrobe item
type Item struct {
	model.Base
	UserID      string
	Name        string
	Description string
	Color       string
	Size        string
	Brand       string
	Price       float64
	Currency    string
	Category    string
	Season      string
	Formal      bool
	Quantity    int
	Starred     bool
	Code        string
	URL         string

	Images  []ItemImage
	ImageID *string
}

// NewItemForm returns an item with values parsed from HTTP form
func NewItemForm(r *http.Request) (*Item, error) {
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		return nil, errors.Annotate(err, "parsing multipart form failed")
	}

	name := strings.TrimSpace(r.FormValue("name"))
	if name == "" {
		return nil, errors.New("please enter a name")
	}

	price, err := form.Float(r, "price")
	if err != nil {
		return nil, errors.Annotate(err, "parsing price failed")
	}

	quantity, err := form.Int(r, "quantity")
	if err != nil {
		return nil, errors.Annotate(err, "parsing quantity failed")
	}

	starred, err := form.Bool(r, "star")
	if err != nil {
		return nil, errors.Annotate(err, "parsing star failed")
	}

	b, err := form.File(r, "image")
	if err != nil {
		return nil, errors.Annotate(err, "parsing image failed")
	}

	var item Item
	item.Name = name
	item.Description = strings.TrimSpace(r.FormValue("description"))
	item.Color = strings.TrimSpace(r.FormValue("color"))
	item.Size = strings.TrimSpace(r.FormValue("size"))
	item.Brand = strings.TrimSpace(r.FormValue("brand"))
	item.Category = strings.TrimSpace(r.FormValue("category"))
	item.Currency = strings.TrimSpace(r.FormValue("currency"))
	item.Season = strings.TrimSpace(r.FormValue("season"))
	item.Price = price
	item.Quantity = quantity
	item.Starred = starred
	item.Code = strings.TrimSpace(r.FormValue("code"))
	item.URL = strings.TrimSpace(r.FormValue("url"))
	item.CreatedAt = time.Now()

	if b != nil {
		itemImage := ItemImage{
			Image: model.Image{
				Body: b,
			},
		}

		item.Images = append(item.Images, itemImage)
	}

	return &item, nil
}
