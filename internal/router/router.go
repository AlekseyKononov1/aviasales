package router

import (
	"log/slog"
	"net/http"

	mhttp "aviasales/internal/http"
	"aviasales/internal/service"

	"github.com/go-chi/chi/v5"
)

func New(svc *service.BookingService, logger *slog.Logger) http.Handler {
	r := chi.NewRouter()

	r.Get("/test", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status": "rly ok"}`))
	})

	h := mhttp.NewHandlers(svc)
	r.Post("/bookings", h.CreateBooking)
	r.Put("/booking/{ticket_no}", h.UpdateBooking)
	r.Delete("/booking/{ticket_no}", h.DeleteBooking)
	r.Get("/flights", h.ListAvailableFlights)
	r.Get("/segments/free", h.ListFreeSeats)

	logger.Info("router initialized")
	return r
}
