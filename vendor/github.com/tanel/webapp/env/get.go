package env

import (
	"os"
)

// Get gets an environment variable
func Get(key string) string {
	return os.Getenv(key)
}
