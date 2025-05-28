#!/bin/bash

KRAKEND_ADDRESS='http://localhost:8080'

QUOTAS_GAME=galaga
QUOTAS_GAME=space_harrier
QUOTAS_GAME=r_type

echo "Making request, showing headers..."

for i in $(seq 1 10);
do
    curl -v \
        -H 'Authorization: bart' \
        -H 'X-Game: sonic_the_hedgehog' \
        "${KRAKEND_ADDRESS}/consume_life"
done
