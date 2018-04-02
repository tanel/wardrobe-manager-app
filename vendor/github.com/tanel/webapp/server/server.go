package server

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/juju/errors"
	"github.com/julienschmidt/httprouter"
	"github.com/tanel/webapp/configuration"
	"github.com/tanel/webapp/db"
	"github.com/tanel/webapp/session"
)

// Serve serves an app
func Serve(name string, r *httprouter.Router) {
	db.Init(name)
	configuration.Init(strings.ToUpper(name))
	session.Init(name)

	if configuration.SharedInstance.LogFile != "" {
		f, err := os.OpenFile(configuration.SharedInstance.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}

		defer func() {
			if err := f.Close(); err != nil {
				log.Println(errors.Annotate(err, "closing log file failed"))
			}
		}()

		log.SetOutput(f)
	}

	port := ":" + configuration.SharedInstance.Port
	log.Println("Server starting at http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, r))
}
