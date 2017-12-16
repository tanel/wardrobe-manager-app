package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/router"
)

const port = ":8080"

func main() {
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	r := router.New()

	log.Println("Server starting at http://localhost" + port)

	log.Fatal(http.ListenAndServe(port, r))
}
