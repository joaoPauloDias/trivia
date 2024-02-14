package main

import (
	"log"
	"os"
)

// NewLogger creates a new logger instance.
func NewLogger() *log.Logger {
	return log.New(os.Stdout, "[api] ", log.LstdFlags)
}
