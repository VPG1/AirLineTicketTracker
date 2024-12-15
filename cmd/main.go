package main

import (
	"AirLineTicketTracker/config"
	"AirLineTicketTracker/internal/bot/telegram"
	"AirLineTicketTracker/internal/services"
	in_memory_storage "AirLineTicketTracker/internal/storage/in-memory_storage"
	"fmt"
)

func main() {
	// load config
	cfg := config.MustLoadConfig("config/config.yaml")

	// init storage
	storage := in_memory_storage.NewStorage()

	// init trackingService
	trackingService := services.NewTrackingService(storage)

	// init bot backend
	bot, err := telegram.New(cfg, trackingService)
	if err != nil {
		fmt.Println(err)
		return
	}

	// backend start
	err = bot.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
}
