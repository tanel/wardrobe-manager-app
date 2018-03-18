package controller

import (
	"github.com/tanel/webapp/configuration"
	"github.com/tanel/webapp/http"
	"golang.org/x/oauth2"
)

// GetFacebookLogin starts FB signup/login
func GetFacebookLogin(request *http.Request) {
	url := configuration.FacebookOAuth2.Facebook().AuthCodeURL("state", oauth2.AccessTypeOffline)
	request.Redirect(url)
}
