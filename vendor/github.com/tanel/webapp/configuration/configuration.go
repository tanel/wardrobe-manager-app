package configuration

import (
	"github.com/tanel/webapp/env"
)

// Configuration represents common app configuration
type Configuration struct {
	SessionSecret  string
	FacebookOAuth2 OAuth2
	Port           string
	LogFile        string
}

var SharedInstance *Configuration

// NewConfiguration returns instance
func Init(prefix string) {
	SharedInstance = &Configuration{
		SessionSecret: env.Required(prefix + "_SESSIONSECRET"),
		FacebookOAuth2: OAuth2{
			ClientID:     env.Required(prefix + "_CLIENTID"),
			ClientSecret: env.Required(prefix + "_CLIENTSECRET"),
			RedirectURL:  env.Required(prefix + "_REDIRECTURL"),
		},
		Port:    env.Required(prefix + "_PORT"),
		LogFile: env.Get(prefix + "_LOGFILE"),
	}
}
