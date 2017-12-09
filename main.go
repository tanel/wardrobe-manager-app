package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/tanel/wardrobe-manager-app/db"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	Serve(":8080")
}
