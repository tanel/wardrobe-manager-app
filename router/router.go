package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-manager-app/controller"
)

// New returns new router instance
func New() *httprouter.Router {
	router := httprouter.New()

	router.GET("/signup", controller.GetSignup)
	router.POST("/signup", controller.PostSignup)

	router.GET("/items/:id", controller.GetItem)
	router.POST("/items/:id", controller.PostItem)
	router.GET("/items", controller.GetItems)
	router.GET("/new", controller.GetItemsNew)
	router.POST("/new", controller.PostItemsNew)
	router.GET("/confirm-delete-item/:id", controller.GetConfirmDeleteItem)
	router.POST("/delete-item/:id", controller.PostDeleteItem)

	router.GET("/item-images/:id", controller.GetItemImage)
	router.GET("/thumbnails/:id", controller.GetItemImageThumbnail)
	router.GET("/confirm-delete-item-image/:id", controller.GetConfirmDeleteItemImage)
	router.POST("/delete-item-image/:id", controller.PostDeleteItemImage)

	router.GET("/weight", controller.GetWeightEntries)
	router.POST("/new-weight", controller.PostNewWeight)
	router.GET("/weights/:id", controller.GetWeight)
	router.POST("/weights/:id", controller.PostWeight)
	router.GET("/confirm-delete-weight/:id", controller.GetConfirmDeleteWeight)
	router.POST("/delete-weight/:id", controller.PostDeleteWeight)

	router.GET("/logout", controller.GetLogout)

	router.GET("/", controller.GetIndex)

	// Serve static files from the ./public directory
	router.NotFound = http.FileServer(http.Dir("public"))

	return router
}
