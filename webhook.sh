#!/bin/bash

curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
     "url": "https://us-central1-uc-internal-sandbox.cloudfunctions.net/bot"
   }' \
   https://api.telegram.org/bot404221654:AAHh87fMJ5_Y_7Bj29anw0H2cNXSxbmp4ig/setWebhook