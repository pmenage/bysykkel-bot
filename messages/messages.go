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

// SendLocationKeyboard asks for the user's current location
func (b Bot) SendLocationKeyboard(update tgbotapi.Update, message string) {
	bot := b.Client
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)

	markup := tgbotapi.ReplyKeyboardMarkup{
		Keyboard: [][]tgbotapi.KeyboardButton{
			[]tgbotapi.KeyboardButton{
				tgbotapi.NewKeyboardButtonLocation("Give location"),
			},
			[]tgbotapi.KeyboardButton{
				tgbotapi.NewKeyboardButton("Cancel"),
			},
		},
		OneTimeKeyboard: true,
	}
	msg.ReplyMarkup = markup
	_, err := bot.Send(msg)
	if err != nil {
		panic(err)
	}
}
