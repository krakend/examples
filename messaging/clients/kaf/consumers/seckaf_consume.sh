#!/bin/bash
# Usage: ./seckaf_consume.sh [orderplacement|portfolioupdates]
cd "$(dirname "$0")" || exit 1
KTOPIC=${1:-orderplacement}
kaf consume \
    ${KTOPIC} \
    --config ./kaf_ssl_conf.yaml \
    --cluster seckafkabroker
