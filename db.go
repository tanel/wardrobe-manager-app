package main

import (
	"database/sql"
	"log"
)

func ConnectDB() {
	connStr := "user=tanel dbname=wardrobe sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

}
