package telegram

import (
	"AirLineTicketTracker/internal/services/tracking_service"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

const (
	commandStart = "start"
	commandHelp  = "help"
	commandTrack = "track"
	commandList  = "list"
	commandStop  = "stop"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case commandHelp:
		return b.handleHelpCommand(message)
	case commandTrack:
		return b.handleTrackCommand(message)
	case commandList:
		return b.handleListCommand(message)
	default:
		return nil
	}

	return nil
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

func (b *Bot) handleTrackCommand(message *tgbotapi.Message) error {
	text := message.Text
	text = strings.TrimPrefix(text, "/track") // удаляем префикс сообщения

	flight, err := b.trackingService.TrackFlight(message.Chat.ID, text)
	if errors.Is(err, tracking_service.IncorrectSearchPhrase) { // некорректная поисковая фраза
		msg := tgbotapi.NewMessage(message.Chat.ID, incorrectSearchPhraseResponse)
		_, err = b.tgBotApi.Send(msg)
		if err != nil {
			return err
		}
		return nil
	} else if errors.Is(err, tracking_service.UserNotRegistered) { // пользователь не вводил команду старт
		msg := tgbotapi.NewMessage(message.Chat.ID, userNotInSystem)

		_, err = b.tgBotApi.Send(msg)
		if err != nil {
			return err
		}

		return nil
	} else if errors.Is(err, tracking_service.FlightAlreadyTracked) {
		formatedDate := flight.DepartureAt.Format("January 2, 2006 15:04 Monday")

		msg := tgbotapi.NewMessage(message.Chat.ID,
			fmt.Sprintf(userAlreadyTrackFlight, flight.Origin, flight.Destination, formatedDate, flight.Price))

		_, err = b.tgBotApi.Send(msg)
		if err != nil {
			return err
		}

		return nil
	} else if err != nil {
		return err
	}

	formatedDate := flight.DepartureAt.Format("January 2, 2006 15:04 Monday")

	msg := tgbotapi.NewMessage(message.Chat.ID,
		fmt.Sprintf(trackResponse, flight.Origin, flight.Destination, formatedDate, flight.Price))

	_, err = b.tgBotApi.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleListCommand(message *tgbotapi.Message) error {
	flights := b.trackingService.GetUserFlight(message.Chat.ID)

	if len(flights) == 0 {
		msg := tgbotapi.NewMessage(message.Chat.ID, NoTrackedFlights)

		_, err := b.tgBotApi.Send(msg)
		if err != nil {
			return err
		}

		return nil
	}

	flightsString := ""
	for i, flight := range flights {
		flightsString += "✈️ " + strconv.Itoa(i+1) + ". " +
			flight.Origin + " → " + flight.Destination +
			" (Текущая цена: " + strconv.Itoa(flight.Price) + "$)" + "\n"
	}

	msg := tgbotapi.NewMessage(message.Chat.ID,
		fmt.Sprintf(listResponse, flightsString))

	_, err := b.tgBotApi.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	log.Println(strconv.Itoa(int(message.Chat.ID)) +
		":" + message.Text)
	return nil
}
