#!/bin/bash

./build-binary.sh
./build-container.sh lyrical-beach-175121
kubectl delete deployment bysykkel
./scripts/deploy.sh lyrical-beach-175121
