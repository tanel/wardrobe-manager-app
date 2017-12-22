package db

import (
	"database/sql"

	"github.com/juju/errors"
)

var db *sql.DB

// Connect connects to PostgreSQL
func Connect() error {
	connStr := "user=wardrobe dbname=wardrobe sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return errors.Annotate(err, "connecting to PostgreSQL failed")
	}

	return nil
}
