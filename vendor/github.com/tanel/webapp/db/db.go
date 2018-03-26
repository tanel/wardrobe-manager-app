package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // import db driver
)

// SharedInstance represents shared DB
var SharedInstance *sql.DB

// Init initializes shared DB
func Init(name string) {
	SharedInstance = Connect(name, name)
}

// Exec executes a query without returning any rows.
func Exec(query string, args ...interface{}) (sql.Result, error) {
	return SharedInstance.Exec(query, args...)
}

// Query executes a query that returns rows, typically a SELECT.
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	return SharedInstance.Query(query, args...)
}

// QueryRow executes a query that is expected to return at most one row.
func QueryRow(query string, args ...interface{}) *sql.Row {
	return SharedInstance.QueryRow(query, args...)
}

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
