package main

import (
	"aviasales/internal/config"
	"aviasales/internal/db"
	"aviasales/internal/logs"
	"aviasales/internal/router"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	cfg := config.Load()

	logger := logs.New(os.Stdout)

	pool, err := db.New(
		db.ConnParams{
			Host:     cfg.Host,
			User:     cfg.User,
			Password: cfg.Password,
			DBName:   cfg.DBName,
		},
		logger,
	)
	if err != nil {
		logger.Error("db init failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer pool.Close()

	r := router.New(logger)

	srv := &http.Server{
		Addr:         cfg.Addr,
		Handler:      r,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	logger.Info("server starting", slog.String("addr", cfg.Addr))
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Error("server error", slog.String("error", err.Error()))
	}
}
