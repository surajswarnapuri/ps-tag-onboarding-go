package main

import (
	"log"

	httpServer "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/interface/http"
)

func main() {
	// Create and start the server
	server := httpServer.NewServer()

	// Start the server (this will block until shutdown)
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
