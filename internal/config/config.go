package config

import (
	"os"
	"time"
)

type Config struct {
	Host         string
	User         string
	Password     string
	DBName       string
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func Load() Config {
	return Config{
		Host:         getenv("DB_HOST", "localhost"),
		User:         getenv("DB_USER", "postgres"),
		Password:     getenv("DB_PASSWORD", "postgres"),
		DBName:       getenv("DB_NAME", "postgres"),
		Addr:         getenv("Addr", ":8080"),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func getenv(key, def string) string {
	v := os.Getenv(key)

	if v != "" {
		return v
	}

	return def
}
