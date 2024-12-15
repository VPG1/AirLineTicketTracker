package aviasales_client

import "time"

type Flight struct {
	Airline      string    `json:"airline"`
	DepartureAt  time.Time `json:"departure_at"`
	ReturnAt     time.Time `json:"return_at"`
	ExpiresAt    time.Time `json:"expires_at"`
	Price        int       `json:"price"`
	FlightNumber int       `json:"flight_number"`
}

type Response struct {
	Data     map[string]map[string]Flight `json:"data"`
	Currency string                       `json:"currency"`
	Success  bool                         `json:"success"`
}
