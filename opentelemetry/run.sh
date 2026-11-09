#!/bin/bash
#

export VERSION=3.0.3
docker pull krakend/krakend-ee:${VERSION}

COLLECTOR_DIR=./config/datadog/collector

if [ -z "$DD_API_KEY" ]; then
    echo -e "\nNo DD_API_KEY found: telemetry won't be exported to Datadog."
    echo -e "Traces are still available in Jaeger at http://localhost:16686"
    echo -e "To also send them to Datadog, set your key and run this script again:\n\n    export DD_API_KEY=\"YOUR_API_KEY\"\n"
    cp ${COLLECTOR_DIR}/collector.no-datadog.yaml ${COLLECTOR_DIR}/collector.yaml
else
    envsubst < ${COLLECTOR_DIR}/collector.yaml.env > ${COLLECTOR_DIR}/collector.yaml
fi

docker compose up -d
