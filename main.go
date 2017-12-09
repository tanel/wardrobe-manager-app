package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	ConnectDB()
	Serve(":8080")
}
