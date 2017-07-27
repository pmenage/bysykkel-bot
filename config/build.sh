#!/bin/bash

go get -u github.com/jteeuwen/go-bindata/...
go-bindata -pkg config -o credentials.go ./config.yaml 