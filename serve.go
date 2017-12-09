package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Serve(port string) {
	router := httprouter.New()
	// Serve static files from the ./public directory
	router.NotFound = http.FileServer(http.Dir("public"))

	log.Println("Server starting at http://localhost" + port)

	log.Fatal(http.ListenAndServe(port, router))

}
