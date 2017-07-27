#!/bin/bash

sh fix-vendor-error.sh

GO_CLOUD_FN_CUSTOM_FLAGS='-ldflags="-X main.projectID='uc-internal-sandbox'"' \
     go-cloud-fn deploy bot -j -s gs://my-bucket-26 \
     --timeout=300 --memory=1024 \