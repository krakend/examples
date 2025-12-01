#!/bin/bash
#

export VERSION=2.12.0
docker pull krakend/krakend-ee:${VERSION}

if [ -z $DD_API_KEY ]; then
    echo -e "\nNo DD_API_KEY found. Set it with:\n\n    export DD_API_KEY="YOUR_API_KEY"\n"
    exit -1
fi

envsubst < ./config/datadog/collector/collector.yaml.env > ./config/datadog/collector/collector.yaml

docker compose up -d 
