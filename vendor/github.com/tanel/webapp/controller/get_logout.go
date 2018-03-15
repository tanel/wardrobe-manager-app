package controller

import (
	"github.com/tanel/webapp/http"
)

// GetLogout logs user out
func GetLogout(request *http.Request) {
	if ok := request.ClearUserID(); !ok {
		return
	}

	request.Redirect("/")
}
