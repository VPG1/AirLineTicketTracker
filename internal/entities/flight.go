package entities

import "time"

type Flight struct {
	OriginIATA      string
	Origin          string
	DestinationIATA string
	Destination     string
	Price           int
	DepartureAt     time.Time
}
