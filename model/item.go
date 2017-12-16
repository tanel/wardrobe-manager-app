package model

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/juju/errors"
)

// Item represents a wardrobe item
type Item struct {
	ID          string
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

	Images []ItemImage
}

func NewItemForm(r *http.Request) (*Item, error) {
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		return nil, errors.Annotate(err, "parsing multipart form failed")
	}

	var item Item

	item.Name = strings.TrimSpace(r.FormValue("name"))
	if item.Name == "" {
		return nil, errors.New("please enter a name")
	}

	s := strings.TrimSpace(r.FormValue("price"))
	if s != "" {
		price, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, errors.New("please enter a valid price or leave it blank")
		}

		item.Price = price
	}

	item.Description = strings.TrimSpace(r.FormValue("description"))
	item.Color = strings.TrimSpace(r.FormValue("color"))
	item.Size = strings.TrimSpace(r.FormValue("size"))
	item.Brand = strings.TrimSpace(r.FormValue("brand"))
	item.Category = strings.TrimSpace(r.FormValue("category"))
	item.Currency = strings.TrimSpace(r.FormValue("currency"))
	item.Season = strings.TrimSpace(r.FormValue("season"))

	file, _, err := r.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		return nil, errors.Annotate(err, "getting form file failed")
	}

	if file != nil {
		defer file.Close()

		b, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, errors.Annotate(err, "reading form file failed")
		}

		item.Images = append(item.Images, ItemImage{
			Body: b,
		})
	}

	return &item, nil
}
