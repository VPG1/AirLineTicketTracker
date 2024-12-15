package entities

type User struct {
	Username string
	ChatId   int64
	Flights  []Flight
}
