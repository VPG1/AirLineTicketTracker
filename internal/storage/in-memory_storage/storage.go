package in_memory_storage

import (
	"AirLineTicketTracker/internal/entities"
	"fmt"
	"log"
)

type Storage struct {
	users map[int64]*entities.User
}

func NewStorage() *Storage {

	return &Storage{make(map[int64]*entities.User)}
}

func (s *Storage) StoreUser(user entities.User) error {
	user.Flights = make([]entities.Flight, 0)
	fmt.Println("Storing users", s.users)
	if _, ok := s.users[user.ChatId]; ok {
		return UserAlreadyExistsError
	} else {
		s.users[user.ChatId] = &user
		log.Println("User stored: ", user)
	}

	return nil
}

func (s *Storage) StoreUserFlight(chatId int64, flight *entities.Flight) error {
	user, ok := s.users[chatId]
	if !ok {
		return UserNotFoundError
	}

	for _, trackedFlight := range user.Flights {
		if trackedFlight.Origin == flight.Origin && trackedFlight.Destination == flight.Destination {
			return FlightAlreadyStored
		}
	}

	s.users[chatId].Flights = append(user.Flights, *flight)

	return nil
}

func (s *Storage) GetUserFlight(chatId int64) []entities.Flight {
	return s.users[chatId].Flights
}
