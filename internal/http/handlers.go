package http

import (
	"encoding/json"
	"net/http"

	"aviasales/internal/errors"
)

type Handlers struct{}

func NewHandlers() *Handlers {
	return &Handlers{}
}

type createBookingReq struct {
	TicketNo string `json:"ticket_no"`
}

type bookingResp struct {
	TicketNo string `json:"ticket_no"`
	Status   string `json:"status"`
}

func (h *Handlers) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var req createBookingReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errors.BadRequest(w, "invalid json")
		return
	}

	resp := bookingResp{
		TicketNo: req.TicketNo,
		Status:   "created",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)
}
