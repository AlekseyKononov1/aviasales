package models

import "time"

type Booking struct {
	TicketNo   string    `json:"ticket_no"`
	TicketSeat string    `json:"ticket_seat"`
	BoughtTime time.Time `json:"ticket_bought_time"`
	Price      int64     `json:"ticket_price"`
}
