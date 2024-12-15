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

const (
	startResponse = `
Привет! 👋 Я бот для отслеживания авиабилетов.
Я помогу тебе найти лучшие цены на рейсы и уведомлю, когда они снизятся.

Вот что я могу:
- ✈️ Отслеживать цены на билеты.
- 🔔 Уведомлять о снижении цен.
- 🕒 Присылать регулярные отчёты.

Напиши /help, чтобы узнать подробнее о командах
`

	helpResponse = `
Вот список доступных команд:  

- /track <город вылета> <город назначения>  
  ➡️ Добавить рейс для отслеживания.  

- /list  
  📋 Показать все рейсы, которые ты отслеживаешь.  

- /remove <номер рейса>  
  🗑 Удалить рейс из списка.  

- /notifications <on|off>  
  🔔 Включить или выключить уведомления.  

- /settings  
  ⚙️ Настройки уведомлений.  

Если у тебя есть вопросы или предложения, напиши нам!
`
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
