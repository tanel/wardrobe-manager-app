package middleware

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	commonhttp "github.com/tanel/webapp/http"
)

// RequestHandlerFunc is a request handler
type RequestHandlerFunc func(request *commonhttp.Request)

func handleRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params, requestHandlerFunc RequestHandlerFunc) {
	logRequest(r)

	request, err := commonhttp.NewRequest(w, r, ps)
	if err != nil {
		log.Println(err)
		http.Error(w, "handling request failed", http.StatusInternalServerError)
		return
	}

	requestHandlerFunc(request)
}
