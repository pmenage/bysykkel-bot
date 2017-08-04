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

	gotra.InitGotra("/translation")

	telegramKey, bysykkelKey := config.GetKeys()
	users := make(messages.Users)
	bot := messages.NewBot(telegramKey)
	//bot.Client.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.Client.GetUpdatesChan(u)
	for update := range updates {

		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		gotra.SetCurrentLanguage(update.Message.From.LanguageCode)
		switch update.Message.From.LanguageCode {
		case "fr-FR":
			gotra.SetCurrentLanguage("Francais")
		default:
			gotra.SetCurrentLanguage("English")
		}

		if _, ok := users[chatID]; !ok {
			users[chatID] = &messages.UserConfig{}
		}

		if users[chatID].Language != "" {
			gotra.SetCurrentLanguage(users[chatID].Language)
		} else if users[chatID].Language == "" &&
			users[chatID].LastMessage != "/language" && users[chatID].LastMessage != "/start" &&
			update.Message.Text != "/language" && update.Message.Text != "/start" {
			bot.SendMessage(update, gotra.T("language.wrong"))
			continue
		}

		if update.Message.Location != nil {

			bot.SendMessage(update, gotra.T("thank"))

			loc := tgbotapi.NewLocation(chatID, update.Message.Location.Latitude, update.Message.Location.Longitude)
			stations := bysykkel.GetStations(bysykkelKey)
			availability := bysykkel.GetStationsAvailability(bysykkelKey)
			status := bysykkel.GetStatus(bysykkelKey)

			if status.Status.AllStationsClosed {
				bot.SendMessage(update, gotra.T("location.allclosed"))
				continue
			}

			msgText := ""
			switch users[chatID].LastMessage {
			case "/bikes":
				bot.SendMessage(update, gotra.T("location.getbikes"))
				msgText = bysykkel.GetNearestBikes(loc.Latitude, loc.Longitude, stations, availability, status,
					gotra.T("location.message"), gotra.T("location.closed"), gotra.T("location.bike"))
			case "/locks":
				bot.SendMessage(update, gotra.T("location.getlocks"))
				msgText = bysykkel.GetNearestLocks(loc.Latitude, loc.Longitude, stations, availability, status,
					gotra.T("location.message"), gotra.T("location.closed"), gotra.T("location.lock"))
			default:
				bot.SendMessage(update, gotra.T("location.retry"))
			}
			bot.SendMessage(update, msgText)
			continue

		}

		log.Printf("[%s] %s", update.Message.From.FirstName, update.Message.Text)

		switch update.Message.Text {
		case "/start", "/language":
			users[chatID].LastMessage = update.Message.Text
			bot.SendLanguageKeyboard(update, gotra.T("language.ask"))
		case "English", "Francais":
			users[chatID].Language = update.Message.Text
			gotra.SetCurrentLanguage(update.Message.Text)
			bot.SendMessage(update, gotra.T("start"))
		case "/locks", "/bikes":
			users[chatID].LastMessage = update.Message.Text
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
