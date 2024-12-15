package services

import (
	"AirLineTicketTracker/internal/entities"
	in_memory_storage "AirLineTicketTracker/internal/storage/in-memory_storage"
	"errors"
)

type Storage interface {
	StoreUser(user entities.User) error
}

type TrackingService struct {
	storage Storage
}

func NewTrackingService(storage Storage) *TrackingService {
	return &TrackingService{storage}
}

func (s *TrackingService) AddUser(username string, chatId int64) error {
	newUser := entities.User{Username: username, ChatId: chatId}
	err := s.storage.StoreUser(newUser)
	if errors.Is(err, in_memory_storage.UserAlreadyExistsError) {
		return nil
	} else if err != nil {
		return err
	}

	return nil
}
