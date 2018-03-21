package env

import (
	"fmt"
	"os"
)

func Required(key string) string {
	value := os.Getenv(key)
	if value == "" {
		fmt.Println("Missing env variable", key)
		os.Exit(1)
	}

	return value
}
