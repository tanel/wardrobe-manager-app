package router

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/webapp/configuration"
	"github.com/tanel/webapp/controller"
	"github.com/tanel/webapp/middleware"
	"github.com/tanel/webapp/session"
)

// New returns new router instance
func New(db *sql.DB, sessionStore *session.Store, oauth2cfg *configuration.OAuth2) *httprouter.Router {
	router := httprouter.New()

	router.GET("/signup", middleware.HandlePublic(db, sessionStore, controller.GetSignup))
	router.POST("/signup", middleware.HandlePublic(db, sessionStore, controller.PostSignup))
	router.GET("/logout", middleware.HandlePublic(db, sessionStore, controller.GetLogout))

	if oauth2cfg != nil {
		router.GET("/facebook-login", middleware.HandleOAuth2(db, sessionStore, *oauth2cfg, controller.GetFacebookLogin))
		router.GET("/facebook-login-completed", middleware.HandleOAuth2(db, sessionStore, *oauth2cfg, controller.GetFacebookLoginCompleted))
	}

	router.GET("/", middleware.HandlePublic(db, sessionStore, controller.GetIndex))

	// Serve static files from the ./public directory
	publicFileServer := http.FileServer(http.Dir("public"))
	router.GET("/public/*filepath", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		r.URL.Path = p.ByName("filepath")
		publicFileServer.ServeHTTP(w, r)
	})

	return router
}
