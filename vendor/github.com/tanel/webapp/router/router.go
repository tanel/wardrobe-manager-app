package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/webapp/controller"
	"github.com/tanel/webapp/middleware"
)

// New returns new router instance
func New() *httprouter.Router {
	router := httprouter.New()

	router.GET("/signup", middleware.HandlePublic(controller.GetSignup))
	router.POST("/signup", middleware.HandlePublic(controller.PostSignup))
	router.GET("/logout", middleware.HandlePublic(controller.GetLogout))

	router.GET("/facebook-login", middleware.HandlePublic(controller.GetFacebookLogin))
	router.GET("/facebook-login-completed", middleware.HandlePublic(controller.GetFacebookLoginCompleted))

	router.GET("/", middleware.HandlePublic(controller.GetIndex))

	// Serve static files from the ./public directory
	publicFileServer := http.FileServer(http.Dir("public"))
	router.GET("/public/*filepath", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		r.URL.Path = p.ByName("filepath")
		publicFileServer.ServeHTTP(w, r)
	})

	return router
}
