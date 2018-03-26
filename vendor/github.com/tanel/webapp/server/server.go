package server

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/webapp/configuration"
	"github.com/tanel/webapp/db"
	"github.com/tanel/webapp/session"
)

func Serve(name string, r *httprouter.Router) {
	db.Init(name)
	configuration.Init(strings.ToUpper(name))
	session.Init(name)

	if configuration.SharedInstance.LogFile != "" {
		f, err := os.OpenFile(configuration.SharedInstance.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}

		defer f.Close()

		log.SetOutput(f)
	}

	port := ":" + configuration.SharedInstance.Port
	log.Println("Server starting at http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, r))
}
