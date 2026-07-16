package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/awoladhossain/turfbook-streaming/config"
	"github.com/awoladhossain/turfbook-streaming/handlers"
	"github.com/gorilla/mux"
)

func initLogger() {
	var handler slog.Handler
	// Choose JSON in production for ELK/Loki log ingestion, Text for clean dev experience
	if os.Getenv("APP_ENV") == "production" {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	}
	slog.SetDefault(slog.New(handler))
}

func main() {
	// Initialize logger
	initLogger()

	// Load configuration
	config.Load()
	slog.Info("Configuration loaded successfully")

	// Verify HLS Output Directory is ready/writable
	if err := os.MkdirAll(config.App.HLSoutputDir, 0755); err != nil {
		slog.Error("Failed to create HLS output directory", "path", config.App.HLSoutputDir, "error", err)
		os.Exit(1)
	}
	slog.Debug("HLS directory verified", "path", config.App.HLSoutputDir)

	// Router setup
	r := mux.NewRouter()
	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

	port := config.App.Port
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Channel to listen for errors during server startup
	serverErrors := make(chan error, 1)

	// Start server in a goroutine
	go func() {
		slog.Info("Streaming server starting", "port", port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- fmt.Errorf("failed to start server: %w", err)
		}
	}()

	// Channel to listen for OS signals (termination/interrupt)
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

	// Block until we receive an error or a shutdown signal
	select {
	case err := <-serverErrors:
		slog.Error("Critical server error", "error", err)
		os.Exit(1)

	case sig := <-shutdownSignal:
		slog.Info("Shutdown signal received, starting graceful shutdown...", "signal", sig.String())

		// Establish context with timeout for graceful exit (e.g. 15 seconds)
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Attempt to gracefully shut down HTTP Server
		if err := srv.Shutdown(ctx); err != nil {
			slog.Error("Graceful shutdown failed, forcing close", "error", err)
			if err := srv.Close(); err != nil {
				slog.Error("Server force close failed", "error", err)
			}
		}

		slog.Info("Graceful shutdown completed. Server stopped.")
	}
}
