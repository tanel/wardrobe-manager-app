package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/http"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	http.Serve(":8080")
}
