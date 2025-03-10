#!/bin/bash

docker run \
    -p 4317:4317 \
    -p 4318:4318 \
    --hostname krakend \
    -v "$(pwd)/collector.yaml:/etc/otelcol-contrib/config.yaml" \
    otel/opentelemetry-collector-contrib

