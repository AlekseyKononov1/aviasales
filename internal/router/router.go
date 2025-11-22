package router

import (
	"log/slog"
	"net/http"

	mhttp "aviasales/internal/http"

	"github.com/go-chi/chi/v5"
)

func New(logger *slog.Logger) http.Handler {
	r := chi.NewRouter()

	r.Get("/test", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status": "rly ok"}`))
	})

	h := mhttp.NewHandlers()
	r.Post("/bookings", h.CreateBooking)

	logger.Info("router initialized")
	return r
}
