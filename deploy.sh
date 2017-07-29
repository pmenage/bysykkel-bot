#!/bin/bash

if [ $# -ne 1 ]
then
  echo "Usage: ./deploy.sh ProjectId"
  exit 1
fi

./build-binary.sh
./build-container.sh $1
kubectl delete deployment bysykkel
./scripts/deploy.sh $1
