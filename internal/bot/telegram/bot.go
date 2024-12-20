package telegram

import (
	"AirLineTicketTracker/config"
	"AirLineTicketTracker/internal/services/tracking_service"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	tgBotApi        *tgbotapi.BotAPI
	trackingService *tracking_service.TrackingService
}

func New(config *config.Config, trackingService *tracking_service.TrackingService) (*Bot, error) {
	bot := &Bot{}

	var err error
	bot.tgBotApi, err = tgbotapi.NewBotAPI(config.Telegram.Token)
	if err != nil {
		return nil, fmt.Errorf("error creating telegram bot: %v", err)
	}

	if config.Env == "debug" {
		bot.tgBotApi.Debug = true
	} else if config.Env == "prod" {
		bot.tgBotApi.Debug = false
	} else {
		return nil, fmt.Errorf("invalid env type: %v", config.Env)
	}

	bot.trackingService = trackingService

	return bot, nil
}

func (b *Bot) Start() error {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.tgBotApi.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			err := b.handleCommand(update.Message)

			if err != nil {
				return fmt.Errorf("error while handling command: %v", err)
			}
		} else {
			err := b.handleMessage(update.Message)

			if err != nil {
				return fmt.Errorf("error while handling message: %v", err)
			}
		}
	}

	return nil
}
