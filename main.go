package main

import (
	"bysykkelBot/bysykkel"
	"bysykkelBot/config"
	"bysykkelBot/messages"
	"errors"
	"log"

	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	telegramKey := ""
	bysykkelKey := ""
	if os.Getenv("DEPLOY_KIND") == "local" {
		config := config.FromYAML("config/config.yaml")
		telegramKey = config.TelegramKey
		bysykkelKey = config.BysykkelKey
	} else if os.Getenv("DEPLOY_KIND") == "cloud" {
		telegramKey = os.Getenv("TELEGRAM_KEY")
		bysykkelKey = os.Getenv("BYSYKKEL_KEY")
	} else {
		log.Println("DEPLOY_KIND is not set")
		panic(errors.New("DEPLOY_KIND is not set"))
	}

	bot := messages.NewBot(telegramKey)

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

			lastMessages = append(lastMessages, lastMessage{
				ChatID:  update.Message.Chat.ID,
				Message: update.Message.Text,
			})

			bot.SendLocationKeyboard(update, "Do you allow the bot to use your current location?")

		} else if update.Message.Text == "Cancel" {

			bot.SendMessage(update, "We need your location to be able to tell you which bikes or locks are close to you. Try again later if you want!")

		} else if update.Message.Location != nil {

			bot.SendMessage(update, "Thank you!\n\nHere are the bikes and locks closest to you:")

			log.Printf("\n\nMessage for location given: %v\n\n", update.Message.Text)

			location := tgbotapi.NewLocation(update.Message.Chat.ID, update.Message.Location.Latitude, update.Message.Location.Longitude)
			stations := bysykkel.GetStations(bysykkelKey)
			availability := bysykkel.GetStationsAvailability(bysykkelKey)

			msgText := ""
			for _, message := range lastMessages {
				if message.Message == "/bikes" && message.ChatID == update.Message.Chat.ID {
					msgText = bysykkel.GetNearestBikes(location.Latitude, location.Longitude, stations, availability)
				} else if message.Message == "/locks" && message.ChatID == update.Message.Chat.ID {
					msgText = bysykkel.GetNearestLocks(location.Latitude, location.Longitude, stations, availability)
				} else if message.Message != "/bikes" && message.Message != "/locks" && message.ChatID == update.Message.Chat.ID {
					msgText = "Please try again, enter the /bikes or the /locks command"
				}
			}

			bot.SendMessage(update, msgText)

		} else if update.Message.Text == "/help" {

			bot.SendMessage(update, "Here are the commands you can send to BysykkelBot:\n\n/bikes - get the bikes closest to you\n/locks - get the locks closest to you")

		} else {

			bot.SendMessage(update, "Sorry, I didn't understand your command. Check out /help if you need to refresh your memory.")

		}

	}

}
