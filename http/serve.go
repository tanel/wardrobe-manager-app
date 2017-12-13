package http

import (
	"log"
	"net/http"
)

func Serve(port string) {
	router := NewRouter()

	log.Println("Server starting at http://localhost" + port)

	log.Fatal(http.ListenAndServe(port, router))
}
