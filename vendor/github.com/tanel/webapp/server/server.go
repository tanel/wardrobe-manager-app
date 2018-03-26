package server

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func Serve(r *httprouter.Router, port string, logFile string) {
	if logFile != "" {
		f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}

		defer f.Close()

		log.SetOutput(f)
	}

	log.Println("Server starting at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
