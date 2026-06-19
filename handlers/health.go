package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "ok",
		"service":   "turfbook-streaming",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	}); err != nil {
		slog.Error("Failed to encode health check response", "error", err)
	}
}
