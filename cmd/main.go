package main

import (
	"AirLineTicketTracker/config"
	"AirLineTicketTracker/internal/bot/telegram"
	"AirLineTicketTracker/internal/notifier"
	"AirLineTicketTracker/internal/services/notification_service"
	"AirLineTicketTracker/internal/services/tracking_service"
	"AirLineTicketTracker/internal/storage/postgres"
	"AirLineTicketTracker/pkg/api_clients/aviasales_client"
	iata_code_definition_api "AirLineTicketTracker/pkg/api_clients/travelpayouts_client"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	//// load config
	cfg := config.MustLoadConfig("config/config.yaml")

	// init storage
	storage, err := postgres.NewStorage(cfg)
	if err != nil {
		log.Println(err)
		return
	}

	IATAService := iata_code_definition_api.New("www.travelpayouts.com")
	AviasalesClient := aviasales_client.New(cfg.FlightsAPI.Host, cfg.FlightsAPI.Path, cfg.FlightsAPI.Token)

	// init newNotifier
	newNotifier, err := notifier.NewNotifier(cfg)

	// init notification service
	noificationService := notification_service.New(newNotifier, AviasalesClient, storage)
	if err != nil {
		return
	}

	// init trackingService
	trackingService := tracking_service.NewTrackingService(storage, IATAService, AviasalesClient, noificationService)

	// init bot backend
	bot, err := telegram.New(cfg, trackingService)
	if err != nil {
		log.Println(err)
		return
	}

	// backend start
	err = bot.Start()
	if err != nil {
		log.Println(err)
		return
	}
}
