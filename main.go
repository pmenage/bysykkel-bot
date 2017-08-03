package main

import (
	"bysykkelBot/bysykkel"
	"bysykkelBot/config"
	"bysykkelBot/messages"
	"log"

	"github.com/cpapidas/Golang-Translator"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	gotra.InitGotra("translation")

	telegramKey, bysykkelKey := config.GetKeys()
	users := make(map[int64]messages.UserConfig)
	bot := messages.NewBot(telegramKey)
	bot.Client.Debug = true
	log.Printf("Authorized on account %s", bot.Client.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.Client.GetUpdatesChan(u)
	for update := range updates {

		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID

		if _, ok := users[chatID]; ok {
			if users[chatID].Language != "" {
				gotra.SetCurrentLanguage(users[chatID].Language)
			}
		}

		if update.Message.Location != nil {

			bot.SendMessage(update, gotra.T("thank"))
			log.Printf("\n\nMessage for location given: %v\n\n", update.Message.Text)

			loc := tgbotapi.NewLocation(update.Message.Chat.ID, update.Message.Location.Latitude, update.Message.Location.Longitude)
			stations := bysykkel.GetStations(bysykkelKey)
			availability := bysykkel.GetStationsAvailability(bysykkelKey)

			msgText := ""
			switch users[chatID].LastMessage {
			case "/bikes":
				bot.SendMessage(update, gotra.T("location.getbikes"))
				msgText = bysykkel.GetNearestBikes(loc.Latitude, loc.Longitude, stations, availability)
			case "/locks":
				bot.SendMessage(update, gotra.T("location.getlocks"))
				msgText = bysykkel.GetNearestLocks(loc.Latitude, loc.Longitude, stations, availability)
			default:
				bot.SendMessage(update, gotra.T("location.retry"))
			}
			bot.SendMessage(update, msgText)
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Text {
		case "/start", "/language":
			bot.SendLanguageKeyboard(update)
			users[chatID] = messages.UserConfig{}
		case "English", "Francais":
			users = messages.SetLanguage(users, chatID, update.Message.Text)
			gotra.SetCurrentLanguage(update.Message.Text)
			bot.SendMessage(update, gotra.T("start"))
		case "/locks", "/bikes":
			users = messages.SetLastMessage(users, chatID, update.Message.Text)
			bot.SendLocationKeyboard(update, gotra.T("location.ask"), gotra.T("location.give"), gotra.T("cmd.cancel"))
		case "Cancel", "Annuler":
			bot.SendMessage(update, gotra.T("cancel"))
		case "/help":
			bot.SendMessage(update, gotra.T("help"))
		default:
			bot.SendMessage(update, gotra.T("else"))
		}

	}

}
