package main

import (
	"AirLineTicketTracker/config"
	"AirLineTicketTracker/pkg/api_clients/aviasales_client"
	trl_client "AirLineTicketTracker/pkg/api_clients/travelpayouts_client"
	"fmt"
)

func main() {
	//// load config
	cfg := config.MustLoadConfig("config/config.yaml")
	//
	//// init storage
	//storage := in_memory_storage.NewStorage()
	//
	//// init trackingService
	//trackingService := services.NewTrackingService(storage)
	//
	//// init bot backend
	//bot, err := telegram.New(cfg, trackingService)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//// backend start
	//err = bot.Start()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	client := trl_client.New("www.travelpayouts.com")
	resp, err := client.GetIATACodes("москва дубай")
	if err != nil {
		return
	}

	fmt.Println(resp)

	aviaClient := aviasales_client.New("api.travelpayouts.com", cfg.FlightsAPI.Token)

	info, err := aviaClient.GetFlightInfo(resp.Origin.IATA, resp.Destination.IATA)
	if err != nil {
		return
	}

	fmt.Println(info)
}
