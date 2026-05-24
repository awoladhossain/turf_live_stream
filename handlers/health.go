package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "ok",
		"service":   "turfbook-streaming",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}
