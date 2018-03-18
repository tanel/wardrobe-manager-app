package controller

import (
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/wardrobe-organizer/ui"
	"github.com/tanel/webapp/http"
)

// GetItemsNew renders new item page
func GetItemsNew(request *http.Request, userID string) {
	category := request.QueryParamByName("category")

	request.Render("items-new", ui.NewItemPage(userID, model.Item{
		Currency: "EUR",
		Quantity: 1,
		Category: category,
	}))
}
