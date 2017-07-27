# Bysykkel bot

## Description

This is a Telegram chatbot written in Go which tells you if there are bicycles or locks in Oslo near you. An app already exists, but this chatbot is useful if you are a frequent user of Telegram and want to know very quickly if there are bikes near you.

For regular users of Oslo Bysykkel, it can be useful to get the nearest stations without having to open the app just to know to which station to go, if you are standing in the middle of two stations for example.

## Run

Install the vendor dependencies with glide. For now, you can build and run from the terminal, a webhook will be implemented later. 

## Notes

The chatbot asks you if you wish to share your location, and then uses the Oslo Bysykkel API to check if there are bicycles or locks near you, according to what your request was.

## Configuration

Add a configuration file in the config folder called config.yaml, with the two following lines:
- telegram_key: YourTelegramBotKey
- bysykkel_key: YourBysykkelAPIKey