package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-organizer/controller"
	"github.com/tanel/webapp/configuration"
	"github.com/tanel/webapp/middleware"
	"github.com/tanel/webapp/router"
)

// New returns new router instance
func New() *httprouter.Router {
	router := router.New()

	configuration.LoggedInPage = "/items"

	router.GET("/items/:id", middleware.RequireUser(controller.GetItem))
	router.POST("/items/:id", middleware.RequireUser(controller.PostItem))
	router.GET("/items", middleware.RequireUser(controller.GetItems))
	router.GET("/new", middleware.RequireUser(controller.GetItemsNew))
	router.POST("/new", middleware.RequireUser(controller.PostItemsNew))
	router.GET("/confirm-delete-item/:id", middleware.RequireUser(controller.GetConfirmDeleteItem))
	router.POST("/delete-item/:id", middleware.RequireUser(controller.PostDeleteItem))

	router.GET("/item-images/:id", middleware.RequireUser(controller.GetItemImage))
	router.GET("/thumbnails/:id", middleware.RequireUser(controller.GetItemImageThumbnail))
	router.GET("/confirm-delete-item-image/:id", middleware.RequireUser(controller.GetConfirmDeleteItemImage))
	router.POST("/delete-item-image/:id", middleware.RequireUser(controller.PostDeleteItemImage))

	router.GET("/outfits", middleware.RequireUser(controller.GetOutfits))
	router.GET("/new-outfit", middleware.RequireUser(controller.GetNewOutfit))
	router.POST("/new-outfit", middleware.RequireUser(controller.PostNewOutfit))
	router.GET("/outfits/:id", middleware.RequireUser(controller.GetOutfit))
	router.POST("/outfits/:id", middleware.RequireUser(controller.PostOutfit))
	router.GET("/confirm-delete-outfit/:id", middleware.RequireUser(controller.GetConfirmDeleteOutfit))
	router.POST("/delete-outfit/:id", middleware.RequireUser(controller.PostDeleteOutfit))
	router.GET("/remove-from-outfit/:id", middleware.RequireUser(controller.GetRemoveFromOutfit))

	router.GET("/weight", middleware.RequireUser(controller.GetWeightEntries))
	router.POST("/new-weight", middleware.RequireUser(controller.PostNewWeight))
	router.GET("/weights/:id", middleware.RequireUser(controller.GetWeight))
	router.POST("/weights/:id", middleware.RequireUser(controller.PostWeight))
	router.GET("/confirm-delete-weight/:id", middleware.RequireUser(controller.GetConfirmDeleteWeight))
	router.POST("/delete-weight/:id", middleware.RequireUser(controller.PostDeleteWeight))

	return router
}
