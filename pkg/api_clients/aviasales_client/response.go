package aviasales_client

import (
	"AirLineTicketTracker/internal/entities"
	"time"
)

type Response struct {
	Data []struct {
		DepartDate       time.Time `json:"depart_date"`
		Origin           string    `json:"origin"`
		Destination      string    `json:"destination"`
		Gate             string    `json:"gate"`
		FoundAt          time.Time `json:"found_at"`
		TripClass        int       `json:"trip_class"`
		Value            int       `json:"value"`
		NumberOfChanges  int       `json:"number_of_changes"`
		Duration         int       `json:"duration"`
		Distance         int       `json:"distance"`
		ShowToAffiliates bool      `json:"show_to_affiliates"`
		Actual           bool      `json:"actual"`
	} `json:"data"`
	Currency string `json:"currency"`
	Success  bool   `json:"success"`
}

func MapRespToFlight(resp *Response, flight *entities.Flight) {
	flight.Price = resp.Data[0].Value
	flight.DepartureAt = resp.Data[0].DepartDate
}
