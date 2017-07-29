package main

import (
	"bysykkelBot/bysykkel"
	"bysykkelBot/messages"
	"log"

	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	//config := config.FromYAML("config/config.yaml")

	bot := messages.NewBot(os.Getenv("TELEGRAM_KEY"))

	bot.Client.Debug = true

	log.Printf("Authorized on account %s", bot.Client.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	type lastMessage struct {
		ChatID  int64
		Message string
	}

	var lastMessages []lastMessage

	updates, _ := bot.Client.GetUpdatesChan(u)

	for update := range updates {

		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.Text == "/start" {

			bot.SendMessage(update, "Hi "+update.Message.Chat.FirstName+", I'm the bysykkel bot, you can send me a message to see if there are bikes or locks near you.\nYou can send the following commands:\n\n/bikes - get the bikes closest to you\n/locks - get the locks closest to you\n/help - see all possible commands\n")

		} else if update.Message.Text == "/locks" || update.Message.Text == "/bikes" {
			lastMsg := lastMessage{
				ChatID:  update.Message.Chat.ID,
				Message: update.Message.Text,
			}
			lastMessages = append(lastMessages, lastMsg)
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
			_, err := bot.Client.Send(msg)
			if err != nil {
				panic(err)
			}

		} else if update.Message.Text == "Cancel" {

			bot.SendMessage(update, "We need your location to be able to tell you which bikes or locks are close to you. Try again later if you want!")

		} else if update.Message.Location != nil {

			bot.SendMessage(update, "Thank you!\n\nHere are the bikes and locks closest to you:")

			log.Printf("\n\nMessage for location given: %v\n\n", update.Message.Text)

			location := tgbotapi.NewLocation(update.Message.Chat.ID, update.Message.Location.Latitude, update.Message.Location.Longitude)
			stations := bysykkel.GetStations(os.Getenv("BYSYKKEL_KEY"))
			availability := bysykkel.GetStationsAvailability(os.Getenv("BYSYKKEL_KEY"))

			msgText := ""
			for _, message := range lastMessages {
				if message.Message == "/bikes" && message.ChatID == update.Message.Chat.ID {
					msgText = bysykkel.GetNearestBikes(location.Latitude, location.Longitude, stations, availability)
				} else if message.Message == "/locks" && message.ChatID == update.Message.Chat.ID {
					msgText = bysykkel.GetNearestLocks(location.Latitude, location.Longitude, stations, availability)
				} else {
					msgText = "We messed up, sorry."
				}
			}

			bot.SendMessage(update, msgText)

		} else if update.Message.Text == "/help" {

			bot.SendMessage(update, "Here are the commands you can send to BysykkelBot:\n\n/bikes - get the bikes closest to you\n/locks - get the locks closest to you")

		} else if update.Message.Text == "/helpmeplease" {

			bot.SendMessage(update, "Coucou mon chéri <3")

		} else if update.Message.Text == "/kisskisslovelove" {

			documentConfig := tgbotapi.NewDocumentUpload(update.Message.Chat.ID, "/paupau.jpg")
			tgbotapi.NewDocumentShare(update.Message.Chat.ID, documentConfig.BaseFile.FileID)

			bot.Client.Send(documentConfig)

			bot.SendMessage(update, "I love you my sweetheart <3")

		} else {

			bot.SendMessage(update, "Sorry, I didn't understand your command. Check out /help if you need to refresh your memory.")

		}

	}

}
