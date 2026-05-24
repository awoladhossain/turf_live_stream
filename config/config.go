package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	RTMPPort     string
	RedisURL     string
	JWTSecret    string
	NestJSURL    string
	HLSoutputDir string
	FFmpegPath   string
}

var App *Config

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

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
