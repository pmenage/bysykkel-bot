# Bysykkel bot

## Description

This is a Telegram chatbot written in Go which tells you if there are bicycles or locks in Oslo near you. An app already exists, but this chatbot is useful if you are a frequent user of Telegram and want to know very quickly if there are bikes near you.

For regular users of Oslo Bysykkel, it can be useful to get the nearest stations without having to open the app just to know to which station to go, if you are standing in the middle of two stations for example.

This bot runs on Google Container Engine, to test it just look for @bysykkelBot on Telegram.

## Prerequisites

You must create a bot with the BotFather on Telegram, and keep the key he sends you to call your bot. You must also create a developer account on Oslo Bysykkel to get a key to access the API.

## Test on local machine

Add a configuration file in the config folder called config.yaml, with the two following lines:
- telegram_key: YourTelegramBotKey
- bysykkel_key: YourBysykkelAPIKey
Set the environment variable DEPLOY\_KIND as "local" or "cloud"

## Deploy on Google Container Engine

Assuming that you have a Google Platform account, gcloud, kubectl, docker, go, and glide installed, and have a GKE cluster to deploy to, follow these steps:
- Install the vendor dependencies with glide with `glide install`
- Setup the secrets with `./scripts/secrets YourTelegramKey YourBysykkelKey`
- Call the deploy script from the bysykkelBot folder `./deploy.sh YourGCPProjectID DeployKind` (DeployKind is "local" or "cloud")

## Notes

The chatbot asks you if you wish to share your location, and then uses the Oslo Bysykkel API to check if there are bicycles or locks near you, according to what your request was.

