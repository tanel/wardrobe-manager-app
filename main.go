package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tanel/wardrobe-organizer/router"
	"github.com/tanel/webapp/db"
	"github.com/tanel/webapp/session"
)

func main() {
	databaseConnection := db.Connect("wardrobe", "wardrobe")

	sessionSecret := os.Getenv("WARDROBE_SESSIONSECRET")
	sessionStore := session.New(sessionSecret, "wardrobe-session")

	r := router.New(databaseConnection, sessionStore)

	port := ":" + os.Getenv("WARDROBE_PORT")

	log.Println("Server starting at http://localhost" + port)

	log.Fatal(http.ListenAndServe(port, r))
}
