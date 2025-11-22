package models

type Segment struct {
	TicketNo      string `json:"ticket_no"`
	FlightID      int    `json:"flight_id"`
	FareCondition string `json:"fare_conditions"`
	Price         string `json:"price"`
}
