package models

import "time"

type Flight struct {
	FlightID                int       `json:"flight_id"`
	ScheduledDepartureLocal time.Time `json:"scheduled_departure_local"`
	ScheduledArrivalLocal   time.Time `json:"scheduled_arrival_local"`
	AirplaneModel           string    `json:"airplane_model"`
	DepartureAirportName    string    `json:"departure_airport_name"`
	DepartureCity           string    `json:"departure_city"`
	DepartureCountry        string    `json:"departure_country"`
	DepartureCoordinates    string    `json:"departure_coordinates"`
	DepartureTimezone       string    `json:"departure_timezone"`
	ArrivalAirportName      string    `json:"arrival_airport_name"`
	ArrivalCity             string    `json:"arrival_city"`
	ArrivalCountry          string    `json:"arrival_country"`
	ArrivalCoordinates      string    `json:"arrival_coordinates"`
	ArrivalTimezone         string    `json:"arrival_timezone"`
}
