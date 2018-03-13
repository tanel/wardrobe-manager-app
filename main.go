package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/tanel/wardrobe-organizer/router"
	"github.com/tanel/webapp/db"
	"github.com/tanel/webapp/session"
)

const port = ":8080"
const unsecureSessionSecret = "7215B41B-812A-4FBF-9A5D-D3ACDE950C12"

func main() {
	databaseConnection := db.Connect("wardrobe", "wardrobe")

	sessionSecret := unsecureSessionSecret
	if secret := os.Getenv("SESSION_SECRET"); secret != "" {
		sessionSecret = secret
	} else {
		log.Println("Warning: set SESSION_SECRET")
	}

	sessionStore := session.New(sessionSecret, "wardrobe-session")

	r := router.New(databaseConnection, sessionStore)

	log.Println("Server starting at http://localhost" + port)

	log.Fatal(http.ListenAndServe(port, r))
}
