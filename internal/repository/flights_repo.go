package repository

import (
	"context"
	"database/sql"
	"time"

	"aviasales/internal/models"
)

type FlightsRepo struct {
	db *sql.DB
}

func NewFlightsRepo(db *sql.DB) *FlightsRepo {
	return &FlightsRepo{db: db}
}

func (r *FlightsRepo) ListAvailable(ctx context.Context, start, end time.Time, depCity, depCountry string, arrCountries []string) ([]models.Flight, error) {
	const q = `
SELECT t.flight_id, 
       t.scheduled_departure_local,
       t.scheduled_arrival_local,
       apl.model AS airplane_model,
       apo_dep.airport_name AS departure_airport_name,
       apo_dep.city AS departure_city,
       apo_dep.country AS departure_country,
       apo_dep.coordinates AS departure_coordinates,
       apo_dep.timezone AS departure_timezone,
       apo_arr.airport_name AS arrival_airport_name,
       apo_arr.city AS arrival_city,
       apo_arr.country AS arrival_country,
       apo_arr.coordinates AS arrival_coordinates,
       apo_arr.timezone AS arrival_timezone
FROM timetable t
INNER JOIN airplanes apl ON t.airplane_code = apl.airplane_code 
INNER JOIN airports apo_dep ON t.departure_airport = apo_dep.airport_code
INNER JOIN airports apo_arr ON t.arrival_airport = apo_arr.airport_code 
WHERE t.scheduled_departure_local >= $1
  AND t.scheduled_arrival_local <= $2
  AND t.status != 'Delayed'
  AND apo_dep.country = $3
  AND apo_dep.city = $4
  AND apo_arr.country = ANY($5);
`
	rows, err := r.db.QueryContext(ctx, q, start, end, depCountry, depCity, pqArray(arrCountries))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flights []models.Flight
	for rows.Next() {
		var f models.Flight
		err := rows.Scan(
			&f.FlightID,
			&f.ScheduledDepartureLocal,
			&f.ScheduledArrivalLocal,
			&f.AirplaneModel,
			&f.DepartureAirportName,
			&f.DepartureCity,
			&f.DepartureCountry,
			&f.DepartureCoordinates,
			&f.DepartureTimezone,
			&f.ArrivalAirportName,
			&f.ArrivalCity,
			&f.ArrivalCountry,
			&f.ArrivalCoordinates,
			&f.ArrivalTimezone,
		)
		if err != nil {
			return nil, err
		}
		flights = append(flights, f)
	}
	return flights, rows.Err()
}

func pqArray(ss []string) string {
	if len(ss) == 0 {
		return "{}"
	}
	out := "{"
	for i, s := range ss {
		if i > 0 {
			out += ","
		}
		out += `"` + s + `"`
	}
	out += "}"
	return out
}
