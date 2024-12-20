package main

import (
	"AirLineTicketTracker/config"
	"AirLineTicketTracker/internal/bot/telegram"
	"AirLineTicketTracker/internal/services"
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

	// init trackingService
	trackingService := services.NewTrackingService(storage, IATAService, AviasalesClient)

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

//package main
//
//import (
//	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
//	"log"
//)
//
//func main() {
//	bot, err := tgbotapi.NewBotAPI("7887255929:AAFJJI7KkC1lJwPKd5kuwX7qn-1ofgmDBYg")
//	if err != nil {
//		log.Panic(err)
//	}
//
//	bot.Debug = true
//
//	log.Printf("Authorized on account %s", bot.Self.UserName)
//
//	u := tgbotapi.NewUpdate(0)
//	u.Timeout = 60
//
//	updates := bot.GetUpdatesChan(u)
//
//	for update := range updates {
//		if update.Message != nil { // If we got a message
//			go foo(update.Message.Chat.ID)
//			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
//
//			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
//			msg.ReplyToMessageID = update.Message.MessageID
//
//			bot.Send(msg)
//		}
//	}
//}
//
//func foo(chatId int64) {
//	bot, err := tgbotapi.NewBotAPI("7887255929:AAFJJI7KkC1lJwPKd5kuwX7qn-1ofgmDBYg")
//	if err != nil {
//		log.Panic(err)
//	}
//
//	bot.Debug = true
//
//	log.Printf("Authorized on account %s", bot.Self.UserName)
//
//	//time.Sleep(10 * time.Second)
//
//	msg := tgbotapi.NewMessage(chatId, "Hello World")
//	if _, err := bot.Send(msg); err != nil {
//		log.Panic(err)
//	}
//}
