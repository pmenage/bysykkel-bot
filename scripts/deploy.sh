#!/bin/bash

set -e

if [ $# -ne 1 ]
then
  echo "Usage: ./deploy.sh ProjectId"
  exit 1
fi

_sh_source="$0"
_dir="$( cd -P "$( dirname "$_sh_source" )" && pwd )"

SHA=$(git rev-parse --short HEAD)

_kube_manifest=$_dir/kube-manifest.yml
cat $_kube_manifest | \
sed -e s/'$TAG'/$SHA/g | \
sed -e s/'$PROJECT'/$1/g | \
kubectl apply -f -