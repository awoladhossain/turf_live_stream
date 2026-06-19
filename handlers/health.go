// Package handlers provides HTTP and WebSocket request handlers
// for the turfbook-streaming service endpoints.
package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

// HealthCheck responds with the current service health status in JSON format.
func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "ok",
		"service":   "turfbook-streaming",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	}); err != nil {
		slog.Error("Failed to encode health check response", "error", err)
	}
}
