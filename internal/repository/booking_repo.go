package repository

import (
	"context"
	"database/sql"

	"aviasales/internal/models"
)

type BookingRepo struct {
	db *sql.DB
}

func NewBookingRepo(db *sql.DB) *BookingRepo {
	return &BookingRepo{db: db}
}

func (r *BookingRepo) CreateFromTicketInfo(ctx context.Context, ticketNo string) (*models.Booking, error) {
	const q = `
WITH ticket_info AS (
    SELECT ticket_no, ticket_seat, ticket_bought_time, ticket_price
    FROM ticket_info
    WHERE ticket_no = $1
)
INSERT INTO bought_tickets (ticket_no, ticket_seat, ticket_bought_time, ticket_price)
SELECT ticket_no, ticket_seat, ticket_bought_time, ticket_price FROM ticket_info
RETURNING ticket_no, ticket_seat, ticket_bought_time, ticket_price;
`
	var b models.Booking
	err := r.db.QueryRowContext(ctx, q, ticketNo).Scan(&b.TicketNo, &b.TicketSeat, &b.BoughtTime, &b.Price)
	return &b, err
}

func (r *BookingRepo) Update(ctx context.Context, ticketNo, seat string, price string) (*models.Booking, error) {
	const q = `
UPDATE bought_tickets
SET ticket_seat = $1,
    ticket_price = $2,
    ticket_bought_time = CURRENT_TIMESTAMP
WHERE ticket_no = $3
RETURNING ticket_no, ticket_seat, ticket_bought_time, ticket_price;
`
	var b models.Booking
	err := r.db.QueryRowContext(ctx, q, seat, price, ticketNo).Scan(&b.TicketNo, &b.TicketSeat, &b.BoughtTime, &b.Price)
	return &b, err
}

func (r *BookingRepo) Delete(ctx context.Context, ticketNo string) error {
	const q = `DELETE FROM bought_tickets WHERE ticket_no = $1;`
	_, err := r.db.ExecContext(ctx, q, ticketNo)
	return err
}
