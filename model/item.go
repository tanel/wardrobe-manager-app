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

func formFloat(r *http.Request, name string) (float64, error) {
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

func formInt(r *http.Request, name string) (int, error) {
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

func formBool(r *http.Request, name string) (bool, error) {
	boolValue, err := strconv.ParseBool(r.FormValue(name))
	if err != nil {
		return false, errors.Annotate(err, "parsing bool failed")
	}

	return boolValue, nil
}

func formFile(r *http.Request, name string) ([]byte, error) {
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

	price, err := formFloat(r, "price")
	if err != nil {
		return nil, errors.Annotate(err, "parsing price failed")
	}

	quantity, err := formInt(r, "quantity")
	if err != nil {
		return nil, errors.Annotate(err, "parsing quantity failed")
	}

	starred, err := formBool(r, "star")
	if err != nil {
		return nil, errors.Annotate(err, "parsing star failed")
	}

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
	item.CreatedAt = time.Now()

	b, err := formFile(r, "image")
	if err != nil {
		return nil, errors.Annotate(err, "parsing image failed")
	}

	item.Images = append(item.Images, ItemImage{
		Body: b,
	})

	return &item, nil
}
