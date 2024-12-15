package main

import (
	"AirLineTicketTracker/config"
	"fmt"
)

func main() {
	// load config
	cfg := config.MustLoadConfig("config/config.yaml")

	fmt.Printf("%+v\n", cfg)
	fmt.Println("Telegram Token: ", cfg.Telegram.Token)
	fmt.Println("Database Host", cfg.Database.Host)
}
