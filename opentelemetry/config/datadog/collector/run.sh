#!/bin/bash

if [ -z $DD_API_KEY ]; then
    echo -e "\nNo DD_API_KEY found. Set it with:\n\n    export DD_API_KEY="YOUR_API_KEY"\n"
    exit -1
fi

# replace the variable in the config file:
envsubst < collector.yaml.env > collector.yaml

docker run \
    -p 4317:4317 \
    -p 4318:4318 \
    --hostname krakend \
    -v "$(pwd)/collector.yaml:/etc/otelcol-contrib/config.yaml" \
    otel/opentelemetry-collector-contrib
