package notifier

import (
	"AirLineTicketTracker/config"
	"AirLineTicketTracker/internal/entities"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const Notification = `
	📉 Отличные новости! Цена на рейс снизилась!  

	📍 Откуда: %s 
	📍 Куда: %s
	🗓 Дата: %s 
	💰 Новая цена: %v
	
	Поторопись, предложение может быстро исчезнуть!
`

type Notifier struct {
	tgBotApi *tgbotapi.BotAPI
}

func NewNotifier(config *config.Config) (*Notifier, error) {
	bot := &Notifier{}

	var err error
	bot.tgBotApi, err = tgbotapi.NewBotAPI(config.Telegram.Token)
	if err != nil {
		return nil, fmt.Errorf("error creating notifier: %v", err)
	}

	if config.Env == "debug" {
		bot.tgBotApi.Debug = true
	} else if config.Env == "prod" {
		bot.tgBotApi.Debug = false
	} else {
		return nil, fmt.Errorf("invalid env type: %v", config.Env)
	}

	return bot, nil
}

func (n *Notifier) Notify(chatId int64, oldPrice int, flight *entities.Flight) error {
	formatedDate := flight.DepartureAt.Format("January 2, 2006 15:04 Monday")

	msg := tgbotapi.NewMessage(chatId,
		fmt.Sprintf(Notification, flight.Origin, flight.Destination, formatedDate, flight.Price))

	_, err := n.tgBotApi.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
