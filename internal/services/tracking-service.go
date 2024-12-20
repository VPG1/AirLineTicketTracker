package services

import (
	"AirLineTicketTracker/internal/entities"
	"AirLineTicketTracker/internal/storage/postgres"
	iata_code_definition_api "AirLineTicketTracker/pkg/api_clients/travelpayouts_client"
	"errors"
)

type Storage interface {
	StoreUser(user entities.User) error
	StoreUserFlight(chatId int64, flight *entities.Flight) error
	GetUserFlight(chatId int64) ([]entities.Flight, error)
}

type IATAConverterService interface {
	GetIATACodes(searchPhrase string) (*entities.Flight, error)
}

type FlightInfoService interface {
	GetFlightInfo(flight entities.Flight) (*entities.Flight, error)
}

type FlightNotificatorService interface {
	TrackFlight(chatId int64, flight *entities.Flight) error
	UntrackFlight(chatId int64, flight *entities.Flight) error
}

type TrackingService struct {
	storage             Storage
	convService         IATAConverterService
	FlightService       FlightInfoService
	notificationService FlightNotificatorService
}

func NewTrackingService(storage Storage, convService IATAConverterService,
	flightService FlightInfoService, notificationService FlightNotificatorService) *TrackingService {
	return &TrackingService{storage, convService, flightService, notificationService}
}

func (s *TrackingService) AddUser(username string, chatId int64) error {
	newUser := entities.User{Username: username, ChatId: chatId}
	err := s.storage.StoreUser(newUser)
	if errors.Is(err, postgres.UserAlreadyExistsError) {
		return nil
	} else if err != nil {
		return err
	}

	return nil
}

func (s *TrackingService) TrackFlight(chatId int64, searchPhrase string) (*entities.Flight, error) {
	flight, err := s.convService.GetIATACodes(searchPhrase)
	if errors.Is(err, iata_code_definition_api.IncorrectResponse) {
		return nil, IncorrectSearchPhrase
	} else if err != nil {
		return nil, err
	}

	flight, err = s.FlightService.GetFlightInfo(*flight)
	if err != nil {
		return nil, err
	}

	err = s.storage.StoreUserFlight(chatId, flight)
	if errors.Is(err, postgres.UserNotFoundError) {
		return nil, UserNotRegistered
	} else if errors.Is(err, postgres.FlightAlreadyStored) {
		return flight, FlightAlreadyTracked
	} else if err != nil {
		return flight, err
	}

	return flight, nil
}

func (s *TrackingService) GetUserFlight(chatId int64) []entities.Flight {
	res, err := s.storage.GetUserFlight(chatId)
	if errors.Is(err, postgres.UserNotFoundError) {
		return nil
	} else if err != nil {
		return nil
	}

	return res
}
