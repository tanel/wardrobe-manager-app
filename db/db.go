package db

import (
	"database/sql"
)

var db *sql.DB

func Connect() error {
	connStr := "user=tanel dbname=wardrobe sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	return nil
}
