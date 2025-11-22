package main

import (
	"aviasales/internal/config"
	"aviasales/internal/db"
	"aviasales/internal/logs"
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

	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status": "rly ok"}`))
	})

	srv := &http.Server{
		Addr:         cfg.Addr,
		Handler:      mux,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	logger.Info("server starting", slog.String("addr", cfg.Addr))
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Error("server error", slog.String("error", err.Error()))
	}
}
