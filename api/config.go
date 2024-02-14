package main

import (
	"fmt"
	"os"
)

const defaultPort = "3000"

const defaultAmount = "3"

// getServerPort retrieves the server port from the environment variable or uses a default.
func getServerPort() string {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = defaultPort
	}
	port = fmt.Sprintf(":%s", port)
	return port
}
