package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

const (
	commandStart = "start"
	commandHelp  = "help"
	commandTrack = "track"
	commandStop  = "stop"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case commandHelp:
		return b.handleHelpCommand(message)
	default:
		return nil
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	err := b.trackingService.AddUser(message.Chat.UserName, message.Chat.ID)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, startResponse)
	_, err = b.tgBotApi.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleHelpCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, helpResponse)
	_, err := b.tgBotApi.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

//func (b *Bot) handleStopCommand(message *tgbotapi.Message) error {
//	msg
//}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	log.Println(strconv.Itoa(int(message.Chat.ID)) + ":" + message.Text)
	return nil
}
