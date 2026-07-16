package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

// Config holds all application configuration values loaded from environment variables.
type Config struct {
	Port         string
	RTMPPort     string
	RedisURL     string
	JWTSecret    string
	NestJSURL    string
	HLSoutputDir string
	FFmpegPath   string
}

// App is the global application configuration instance, initialized by Load.
var App *Config

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// Load reads environment variables from .env file (if present) and
// populates the global App config. Falls back to defaults when variables are not set.
func Load() {
	if err := godotenv.Load(); err != nil {
		log.Println("NO .env file found-using evironment variable")
	}

	App = &Config{
		Port:         getEnv("PORT", "8080"),
		RTMPPort:     getEnv("RTMP_PORT", "1935"),
		RedisURL:     getEnv("REDIS_URL", "redis://localhost:6379"),
		JWTSecret:    getEnv("JWT_SECRET", "your_jwt_secret"),
		NestJSURL:    getEnv("NESTJS_URL", "http://localhost:3000"),
		HLSoutputDir: getEnv("HLS_OUTPUT_DIR", "./hls"),
		FFmpegPath:   getEnv("FFMPEG_PATH", "ffmpeg"),
	}
	if App.JWTSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}
}
