package db

import (
	"database/sql"
	"fmt"
	"log"
)

// Connect returns connection or panics
func Connect(userName, dbName string) *sql.DB {
	if userName == "" {
		log.Panic("db username is mandatory")
	}

	if dbName == "" {
		log.Panic("db name is mandatory")
	}

	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable", userName, dbName)
	connection, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panic(err)
	}

	return connection
}
