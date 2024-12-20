package notification_service

import (
	"AirLineTicketTracker/internal/entities"
	"fmt"
	"log"
	"time"
)

type Notifier interface {
	Notify(chatId int64, oldPrice int, flight *entities.Flight) error
}

type FlightInfoService interface {
	GetFlightInfo(flight entities.Flight) (*entities.Flight, error)
}

type Storage interface {
	GetFlightId(chatId int64, flight *entities.Flight) (int64, error)
	GetFlightById(id int64) (*entities.Flight, error)
	ChangeFlightPrice(id int64, price int) error
	GetAllFlights() ([]entities.Flight, error)
	GetUsers() ([]entities.User, error)
}

type NotificationService struct {
	quits      map[int64]chan struct{}
	Notifier   Notifier
	FlightInfo FlightInfoService
	Storage    Storage
}

func New(notifier Notifier, flightInfo FlightInfoService, storage Storage) *NotificationService {
	nf := &NotificationService{make(map[int64]chan struct{}), notifier, flightInfo, storage}

	users, err := storage.GetUsers()
	if err != nil {
		return nil
	}

	for _, user := range users {
		for _, flight := range user.Flights {
			err := nf.TrackFlight(user.ChatId, &flight)
			if err != nil {
				return nil
			}
		}
	}

	return nf
}

func (ns *NotificationService) TrackFlight(chatId int64, flight *entities.Flight) error {
	ticker := time.NewTicker(1 * time.Minute)
	quit := make(chan struct{})

	id, err := ns.Storage.GetFlightId(chatId, flight)
	if err != nil {
		return err
	}

	ns.quits[id] = quit

	go func() {
		fmt.Println(chatId)
		for {
			select {
			case <-ticker.C:
				log.Println("Start tracking", chatId, flight)

				oldFlight, err := ns.Storage.GetFlightById(id)
				if err != nil {
					return
				}

				newFlight, err := ns.FlightInfo.GetFlightInfo(*flight)
				if err != nil {
					return
				}

				if newFlight.Price < oldFlight.Price {
					err := ns.Notifier.Notify(chatId, oldFlight.Price, newFlight)
					if err != nil {
						return
					}
				}

				err = ns.Storage.ChangeFlightPrice(id, newFlight.Price)
				if err != nil {
					return
				}

			case <-quit:
				return
			}
		}
	}()

	return nil
}

func (ns *NotificationService) UntrackFlight(chatId int64, flight *entities.Flight) error {
	//quit := ns.quits[FlightID{chatId, flightId}]
	//quit <- struct{}{}
	//
	//close(ns.quits[FlightID{chatId, flightId}])

	return nil
}
