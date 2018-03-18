package configuration

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

// OAuth2 represents OAuth2 configuration
type OAuth2 struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

// FacebookOAuth2 specifies Facebook OAuth2 configuration
var FacebookOAuth2 OAuth2

// Facebook returns Facebook OAuth2 config
func (cfg OAuth2) Facebook() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Scopes:       []string{"public_profile", "email"},
		RedirectURL:  cfg.RedirectURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  facebook.Endpoint.AuthURL,
			TokenURL: facebook.Endpoint.TokenURL,
		},
	}
}
