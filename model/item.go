package model

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/juju/errors"
)

// Item represents a wardrobe item
type Item struct {
	Base
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

	Images  []ItemImage
	ImageID *string
}

// NewItemForm returns an item with values parsed from HTTP form
func NewItemForm(r *http.Request) (*Item, error) {
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		return nil, errors.Annotate(err, "parsing multipart form failed")
	}

	var item Item

	item.Name = strings.TrimSpace(r.FormValue("name"))
	if item.Name == "" {
		return nil, errors.New("please enter a name")
	}

	if s := strings.TrimSpace(r.FormValue("price")); s != "" {
		s = strings.Replace(s, ",", ".", -1)
		price, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, errors.New("please enter a valid price or leave it blank")
		}

		item.Price = price
	}

	if s := strings.TrimSpace(r.FormValue("quantity")); s != "" {
		quantity, err := strconv.Atoi(s)
		if err != nil {
			return nil, errors.New("please enter a valid quantity or leave it blank")
		}

		item.Quantity = quantity
	}

	item.Description = strings.TrimSpace(r.FormValue("description"))
	item.Color = strings.TrimSpace(r.FormValue("color"))
	item.Size = strings.TrimSpace(r.FormValue("size"))
	item.Brand = strings.TrimSpace(r.FormValue("brand"))
	item.Category = strings.TrimSpace(r.FormValue("category"))
	item.Currency = strings.TrimSpace(r.FormValue("currency"))
	item.Season = strings.TrimSpace(r.FormValue("season"))
	item.CreatedAt = time.Now()

	starred, err := strconv.ParseBool(r.FormValue("star"))
	if err != nil {
		return nil, errors.Annotate(err, "parsing star failed")
	}

	item.Starred = starred

	file, _, err := r.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		return nil, errors.Annotate(err, "getting form file failed")
	}

	if file != nil {
		defer func() {
			if err := file.Close(); err != nil {
				log.Println(errors.Annotate(err, "closing file failed"))
			}
		}()

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
