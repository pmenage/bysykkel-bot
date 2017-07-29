#!/bin/bash

set -e

if [ $# -ne 1 ]
then
    echo "Usage: ./build.sh ProjectId"
    exit 1
fi

SHA=$(git rev-parse --short HEAD)
API_CONTAINER_TAG="gcr.io/$1/bysykkel:$SHA"
docker build -t $API_CONTAINER_TAG .;
gcloud docker -- push $API_CONTAINER_TAG;