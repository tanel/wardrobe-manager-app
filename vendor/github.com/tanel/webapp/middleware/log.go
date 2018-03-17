package middleware

import (
	"log"
	"net/http"
)

func logRequest(r *http.Request) {
	log.Println(r.Method, r.URL)
}
