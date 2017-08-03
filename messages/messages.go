package messages

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// Bot is the Telegram Bot API
type Bot struct {
	Client *tgbotapi.BotAPI
}

// UserConfig contains language and last message
type UserConfig struct {
	Language    string
	LastMessage string
}

// Users contains is a map on the Chat ID
type Users map[int64]*UserConfig

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

// SendLanguageKeyboard sends a keyboard to chose a language
func (b Bot) SendLanguageKeyboard(update tgbotapi.Update) {
	bot := b.Client
	msg := tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"Which language do you speak?")

	markup := tgbotapi.ReplyKeyboardMarkup{
		Keyboard: [][]tgbotapi.KeyboardButton{
			[]tgbotapi.KeyboardButton{
				tgbotapi.NewKeyboardButton("English"),
			},
			[]tgbotapi.KeyboardButton{
				tgbotapi.NewKeyboardButton("Francais"),
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

// SendLocationKeyboard asks for the user's current location
func (b Bot) SendLocationKeyboard(update tgbotapi.Update, message, location, cancel string) {
	bot := b.Client
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)

	markup := tgbotapi.ReplyKeyboardMarkup{
		Keyboard: [][]tgbotapi.KeyboardButton{
			[]tgbotapi.KeyboardButton{
				tgbotapi.NewKeyboardButtonLocation(location),
			},
			[]tgbotapi.KeyboardButton{
				tgbotapi.NewKeyboardButton(cancel),
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
