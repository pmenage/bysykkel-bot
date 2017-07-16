package main

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("404221654:AAHh87fMJ5_Y_7Bj29anw0H2cNXSxbmp4ig")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.Text == "/start" {
			msg := tgbotapi.NewMessage(
				update.Message.Chat.ID,
				"Hi, I'm the bysykkel bot, you can send me a message to see if there are bikes or locks near you.\n You can send the following commands:\n\n /getbikes - get the bikes closest to you\n /getlocks - get the locks closest to you\n")
			msg.ReplyToMessageID = update.Message.MessageID
			_, err := bot.Send(msg)
			if err != nil {
				panic(err)
			}
		}

		if update.Message.Text == "/getbikes" || update.Message.Text == "/getlocks" {
			msg := tgbotapi.NewMessage(
				update.Message.Chat.ID,
				"Do you allow the bot to use your current location?")
			var keyboardButtons []tgbotapi.KeyboardButton
			locationButton := tgbotapi.NewKeyboardButtonLocation("Yes, I allow BysykkelBot to use my location")
			keyboardButtons = append(keyboardButtons, locationButton)
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(keyboardButtons)
			_, err := bot.Send(msg)
			if err != nil {
				panic(err)
			}
		}

	}

}
