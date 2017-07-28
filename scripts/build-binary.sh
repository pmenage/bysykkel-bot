#!/bin/bash

CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -ldflags '-s -w' -installsuffix cgo -o bysykkel main.go