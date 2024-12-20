package postgres

import (
	"AirLineTicketTracker/internal/entities"
	"time"
)

type UsersSchema struct {
	Id       int64  `db:"id"`
	ChatId   int64  `db:"chat_id"`
	Username string `db:"username"`
}

type FlightsSchema struct {
	Id              int64     `db:"id"`
	OriginIATA      string    `db:"origin_iata"`
	Origin          string    `db:"origin"`
	DestinationIATA string    `db:"destination_iata"`
	Destination     string    `db:"destination"`
	Price           int       `db:"price"`
	DepartureAt     time.Time `db:"departure_at"`
	UserId          string    `db:"user_id"`
}

func (f *FlightsSchema) ToFlight() entities.Flight {
	return entities.Flight{OriginIATA: f.OriginIATA, Origin: f.Origin,
		DestinationIATA: f.DestinationIATA, Destination: f.Destination,
		Price: f.Price, DepartureAt: f.DepartureAt}
}
