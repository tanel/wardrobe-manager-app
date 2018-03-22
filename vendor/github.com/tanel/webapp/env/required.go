package env

import (
	"fmt"
	"os"
)

// Required requires an environment variable
func Required(key string) string {
	value := os.Getenv(key)
	if value == "" {
		fmt.Println("Missing env variable", key)
		os.Exit(1)
	}

	return value
}
