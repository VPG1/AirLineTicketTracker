package iata_code_definition_api

import "AirLineTicketTracker/internal/entities"

type Response struct {
	Origin struct {
		IATA string `json:"iata"`
		Name string `json:"name"`
	} `json:"origin"`
	Destination struct {
		IATA string `json:"iata"`
		Name string `json:"name"`
	} `json:"destination"`
}

func getFlightFromResp(response *Response) *entities.Flight {
	flight := entities.Flight{}
	flight.OriginIATA = response.Origin.IATA
	flight.Origin = response.Origin.Name
	flight.DestinationIATA = response.Destination.IATA
	flight.Destination = response.Destination.Name

	return &flight
}
