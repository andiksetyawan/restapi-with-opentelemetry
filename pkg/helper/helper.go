package helper

import (
	"os"
)

//GetEnv with fallback value
func GetEnv(key, fallback string) string {
	if env := os.Getenv(key); env != "" {
		return env
	}
	return fallback
}
