package main

import (
	"github.com/tanel/wardrobe-organizer/router"
	"github.com/tanel/webapp/configuration"
	"github.com/tanel/webapp/db"
	"github.com/tanel/webapp/env"
	"github.com/tanel/webapp/server"
	"github.com/tanel/webapp/session"
)

func main() {
	databaseConnection := db.Connect("wardrobe", "wardrobe")

	sessionSecret := env.Required("WARDROBE_SESSIONSECRET")
	sessionStore := session.New(sessionSecret, "wardrobe-session")

	configuration.FacebookOAuth2.ClientID = env.Required("WARDROBE_CLIENTID")
	configuration.FacebookOAuth2.ClientSecret = env.Required("WARDROBE_CLIENTSECRET")
	configuration.FacebookOAuth2.RedirectURL = env.Required("WARDROBE_REDIRECTURL")

	r := router.New(databaseConnection, sessionStore)
	port := env.Required("WARDROBE_PORT")

	server.Serve(r, port)
}
