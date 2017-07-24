# Bysykkel bot

## Description

This is a Telegram chatbot written in Go which tells you if there are bicycles or locks in Oslo near you.

## Run

For now, you can build and run from the terminal, a webhook will be implemented later.

## Notes

The chatbot asks you if you wish to share your location, and then uses the Oslo Bysykkel API to check if there are bicycles or locks near you, according to what your request was.

## Configuration

Add a configuration file in the config folder called config.yaml, with the two following lines:
- telegram_key: YourTelegramBotKey
- bysykkel_key: YourBysykkelAPIKey