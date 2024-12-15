package in_memory_storage

import (
	"AirLineTicketTracker/internal/entities"
	"fmt"
	"log"
)

type Storage struct {
	users map[int64]entities.User
}

func NewStorage() *Storage {
	return &Storage{make(map[int64]entities.User)}
}

func (s *Storage) StoreUser(user entities.User) error {
	fmt.Println("Storing users", s.users)
	if _, ok := s.users[user.ChatId]; ok {
		return UserAlreadyExistsError
	} else {
		s.users[user.ChatId] = user
		log.Println("User stored: ", user)
	}

	return nil
}
