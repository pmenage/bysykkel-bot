package messages

import "github.com/go-telegram-bot-api/telegram-bot-api"

// Bot is the Telegram Bot API
type Bot struct {
	Client *tgbotapi.BotAPI
}

// NewBot creates a new bot
func NewBot(telegramKey string) Bot {
	bot, err := tgbotapi.NewBotAPI(telegramKey)
	if err != nil {
		panic(err)
	}
	return Bot{
		Client: bot,
	}
}

// SendMessage sends a message to user
func (b Bot) SendMessage(update tgbotapi.Update, message string) {
	bot := b.Client
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	_, err := bot.Send(msg)
	if err != nil {
		panic(err)
	}
}
