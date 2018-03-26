package server

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Serve(r *httprouter.Router, port string) {
	log.Println("Server starting at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
