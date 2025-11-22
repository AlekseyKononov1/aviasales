package db

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

type ConnParams struct {
	Host     string
	User     string
	Password string
	DBName   string
}

func New(p ConnParams, logger *slog.Logger) (*DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		p.Host, p.User, p.Password, p.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
			defer cancel()

			err = db.Ping()
			if err == nil {
				return
			}

			logger.Warn("database not ready, retrying...", slog.String("error", err.Error()))
			select {
			case <-ctx.Done():
				continue
			}

		}
	}()
	wg.Wait()
	logger.Info("database connection established")

	return &DB{DB: db}, nil
}
