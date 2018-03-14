package router

import (
	"database/sql"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-organizer/controller"
	"github.com/tanel/webapp/middleware"
	"github.com/tanel/webapp/router"
	"github.com/tanel/webapp/session"
)

// New returns new router instance
func New(db *sql.DB, sessionStore *session.Store) *httprouter.Router {
	router := router.New(db, sessionStore)

	router.GET("/items/:id", middleware.RequireUser(db, sessionStore, controller.GetItem))
	router.POST("/items/:id", middleware.RequireUser(db, sessionStore, controller.PostItem))
	router.GET("/items", middleware.RequireUser(db, sessionStore, controller.GetItems))
	router.GET("/new", middleware.RequireUser(db, sessionStore, controller.GetItemsNew))
	router.POST("/new", middleware.RequireUser(db, sessionStore, controller.PostItemsNew))
	router.GET("/confirm-delete-item/:id", middleware.RequireUser(db, sessionStore, controller.GetConfirmDeleteItem))
	router.POST("/delete-item/:id", middleware.RequireUser(db, sessionStore, controller.PostDeleteItem))

	router.GET("/item-images/:id", middleware.RequireUser(db, sessionStore, controller.GetItemImage))
	router.GET("/thumbnails/:id", middleware.RequireUser(db, sessionStore, controller.GetItemImageThumbnail))
	router.GET("/confirm-delete-item-image/:id", middleware.RequireUser(db, sessionStore, controller.GetConfirmDeleteItemImage))
	router.POST("/delete-item-image/:id", middleware.RequireUser(db, sessionStore, controller.PostDeleteItemImage))

	router.GET("/outfits", middleware.RequireUser(db, sessionStore, controller.GetOutfits))
	router.GET("/new-outfit", middleware.RequireUser(db, sessionStore, controller.GetNewOutfit))
	router.POST("/new-outfit", middleware.RequireUser(db, sessionStore, controller.PostNewOutfit))
	router.GET("/outfits/:id", middleware.RequireUser(db, sessionStore, controller.GetOutfit))
	router.POST("/outfits/:id", middleware.RequireUser(db, sessionStore, controller.PostOutfit))
	router.GET("/confirm-delete-outfit/:id", middleware.RequireUser(db, sessionStore, controller.GetConfirmDeleteOutfit))
	router.POST("/delete-outfit/:id", middleware.RequireUser(db, sessionStore, controller.PostDeleteOutfit))
	router.GET("/remove-from-outfit/:id", middleware.RequireUser(db, sessionStore, controller.GetRemoveFromOutfit))

	router.GET("/weight", middleware.RequireUser(db, sessionStore, controller.GetWeightEntries))
	router.POST("/new-weight", middleware.RequireUser(db, sessionStore, controller.PostNewWeight))
	router.GET("/weights/:id", middleware.RequireUser(db, sessionStore, controller.GetWeight))
	router.POST("/weights/:id", middleware.RequireUser(db, sessionStore, controller.PostWeight))
	router.GET("/confirm-delete-weight/:id", middleware.RequireUser(db, sessionStore, controller.GetConfirmDeleteWeight))
	router.POST("/delete-weight/:id", middleware.RequireUser(db, sessionStore, controller.PostDeleteWeight))

	return router
}
