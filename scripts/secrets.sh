#!/bin/bash

if [ $# -ne 2 ]
then
    echo "Usage: ./secrets.sh TelegramKey BysykkelKey"
    exit 1
fi

kubectl create secret generic keys --from-literal=telegramkey=$1 --from-literal=bysykkelkey=$2