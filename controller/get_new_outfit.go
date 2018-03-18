package controller

import (
	"github.com/tanel/wardrobe-organizer/model"
	"github.com/tanel/wardrobe-organizer/ui"
	"github.com/tanel/webapp/http"
)

// GetNewOutfit renders new outfit page
func GetNewOutfit(request *http.Request, userID string) {
	request.Render("new-outfit", ui.NewOutfitPage(userID, model.Outfit{}))
}
