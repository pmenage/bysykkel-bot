package main

import (
	"bysykkelBot/bysykkel"
	"bysykkelBot/config"
	"log"
	"net/http"

	"fmt"

	"encoding/json"
	"io/ioutil"

	"github.com/MartinSahlen/go-cloud-fn/shim"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func handler(w http.ResponseWriter, r *http.Request) {

	file, err := config.Asset("config.yaml")
	if err != nil {
		panic(err)
	}
	config := config.FromYAML(file)

	bot, err := tgbotapi.NewBotAPI(config.TelegramKey)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	type lastMessage struct {
		ChatID  int64
		Message string
	}

	var lastMessages []lastMessage

	ch := make(chan tgbotapi.Update, bot.Buffer)

	bytes, _ := ioutil.ReadAll(r.Body)
	var update tgbotapi.Update
	json.Unmarshal(bytes, &update)
	fmt.Println(update)

	ch <- update
	fmt.Printf("ch is: %v", ch)
	// updates, _ := bot.GetUpdatesChan(u)

	for update := range ch {

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
				"Thank you!\n\n Here are the bikes and locks closest to you:")

			_, err = bot.Send(msg)
			if err != nil {
				panic(err)
			}

			log.Printf("\n\nMessage for location given: %v\n\n", update.Message.Text)

			location := tgbotapi.NewLocation(update.Message.Chat.ID, update.Message.Location.Latitude, update.Message.Location.Longitude)
			stations := bysykkel.GetStations(config.BysykkelKey)
			availability := bysykkel.GetStationsAvailability(config.BysykkelKey)

			msgText := ""
			for _, message := range lastMessages {
				if message.Message == "/getbikes" && message.ChatID == update.Message.Chat.ID {
					msgText = bysykkel.GetNearestBikes(location.Latitude, location.Longitude, stations, availability)
				} else if message.Message == "/getlocks" && message.ChatID == update.Message.Chat.ID {
					msgText = bysykkel.GetNearestLocks(location.Latitude, location.Longitude, stations, availability)
				} else {
					msgText = "We messed up, sorry."
				}
			}

			msg = tgbotapi.NewMessage(
				update.Message.Chat.ID,
				msgText)
			_, err = bot.Send(msg)
			if err != nil {
				panic(err)
			}
		}

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func main() {

	shim.ServeHTTP(handler)

}
