package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/awoladhossain/turfbook-streaming/config"
	"github.com/awoladhossain/turfbook-streaming/handlers"
	"github.com/gorilla/mux"
)

func main() {
	config.Load()

	r := mux.NewRouter()

	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

	port := config.App.Port
	fmt.Printf("Streaming server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// $(go env GOPATH)/bin/air

