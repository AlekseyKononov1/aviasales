package repository

import (
	"context"
	"database/sql"

	"aviasales/internal/models"
)

type SegmentsRepo struct {
	db *sql.DB
}

func NewSegmentsRepo(db *sql.DB) *SegmentsRepo {
	return &SegmentsRepo{db: db}
}

func (r *SegmentsRepo) ListFree(ctx context.Context, flightID int, fare string) ([]models.Segment, error) {
	const q = `
SELECT s.ticket_no, s.price
FROM segments s 
WHERE s.flight_id = $1
  AND s.fare_conditions = $2
ORDER BY s.price;
`
	rows, err := r.db.QueryContext(ctx, q, flightID, fare)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []models.Segment
	for rows.Next() {
		var s models.Segment
		err := rows.Scan(&s.TicketNo, &s.Price)
		if err != nil {
			return nil, err
		}
		s.FlightID = flightID
		s.FareCondition = fare
		out = append(out, s)
	}
	return out, rows.Err()
}
