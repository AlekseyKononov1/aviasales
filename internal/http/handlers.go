package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"aviasales/internal/errors"
	"aviasales/internal/service"

	"github.com/go-chi/chi/v5"
)

type Handlers struct {
	svc *service.BookingService
}

func NewHandlers(svc *service.BookingService) *Handlers {
	return &Handlers{svc: svc}
}

func (h *Handlers) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var req struct {
		TicketNo string `json:"ticket_no"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errors.BadRequest(w, "invalid json")
		return
	}

	b, err := h.svc.CreateBooking(r.Context(), req.TicketNo)
	if err != nil {
		errors.Internal(w, err.Error())
		return
	}

	writeJSON(w, http.StatusAccepted, b)
}

func (h *Handlers) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	ticketNo := chi.URLParam(r, "ticket_no")

	var req struct {
		TicketSeat string `json:"ticket_seat"`
		Price      string `json:"ticket_price"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errors.BadRequest(w, "invalid json")
		return
	}

	b, err := h.svc.UpdateBooking(r.Context(), ticketNo, req.TicketSeat, req.Price)
	if err != nil {
		errors.Internal(w, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, b)
}

func (h *Handlers) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	ticketNo := chi.URLParam(r, "ticket_no")
	err := h.svc.DeleteBooking(r.Context(), ticketNo)
	if err != nil {
		errors.Internal(w, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"deleted": ticketNo})
}

func (h *Handlers) ListAvailableFlights(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	startStr := q.Get("start")
	endStr := q.Get("end")
	depCity := q.Get("dep_city")
	depCountry := q.Get("dep_country")
	arrCountries := q["arr_country"]

	start, _ := time.Parse(time.RFC3339, startStr)
	end, _ := time.Parse(time.RFC3339, endStr)

	flights, err := h.svc.ListAvailableFlights(r.Context(), start, end, depCity, depCountry, arrCountries)
	if err != nil {
		errors.Internal(w, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, flights)
}

func (h *Handlers) ListFreeSeats(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	flightIDStr := q.Get("flight_id")
	fare := q.Get("fare_conditions")

	flightID, err := strconv.Atoi(flightIDStr)
	if err != nil {
		errors.BadRequest(w, "invalid flight_id")
		return
	}

	segs, err := h.svc.ListFreeSeats(r.Context(), flightID, fare)
	if err != nil {
		errors.Internal(w, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, segs)
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
