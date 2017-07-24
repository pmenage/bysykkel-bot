package main

import (
	"bysykkelBot/bysykkel"
	"bysykkelBot/config"
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	config := config.FromYAML("config/config.yaml")

	bot, err := tgbotapi.NewBotAPI(config.TelegramKey)
	if err != nil {
		panic(err)
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
			_, err := bot.Send(msg)
			if err != nil {
				panic(err)
			}
		}

		if update.Message.Text == "/getlocks" || update.Message.Text == "/getbikes" {
			msg := tgbotapi.NewMessage(
				update.Message.Chat.ID,
				"Do you allow the bot to use your current location?")
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

		if update.Message.Text == "Cancel" {
			msg := tgbotapi.NewMessage(
				update.Message.Chat.ID,
				"We need your location to be able to tell you which bikes or locks are close to you. Try again later if you want!")
			_, err := bot.Send(msg)
			if err != nil {
				panic(err)
			}
		}

		if update.Message.Location != nil {
			msg := tgbotapi.NewMessage(
				update.Message.Chat.ID,
				"Thank you!")

			_, err = bot.Send(msg)
			if err != nil {
				panic(err)
			}

			location := tgbotapi.NewLocation(update.Message.Chat.ID, update.Message.Location.Latitude, update.Message.Location.Longitude)
			stations := bysykkel.GetStations(config.BysykkelKey)
			availability := bysykkel.GetStationsAvailability(config.BysykkelKey)

			msgText := bysykkel.GetNearestBikes(location.Latitude, location.Longitude, stations, availability)

			msg = tgbotapi.NewMessage(
				update.Message.Chat.ID,
				msgText)
			_, err = bot.Send(msg)
			if err != nil {
				panic(err)
			}
		}

	}

}
