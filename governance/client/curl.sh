#!/bin/bash

KRAKEND_ADDRESS='http://localhost:8080'

# Number of requests to make
NUM_REQUESTS=1

# Endpoint to call
# QUOTAS_ENDPOINT='/request_weapon_power_up'
QUOTAS_ENDPOINT='/request_squad_call'
# QUOTAS_ENDPOINT='/consume_life'

# Player to use
QUOTAS_USER=bart
# QOTAS_USER=homer

# Select the game to use;
QUOTAS_GAME=galaga
# QUOTAS_GAME=space_harrier
# QUOTAS_GAME=r_type

echo "Making ${NUM_REQUESTS} request(s), showing headers..."

for i in $(seq 1 ${NUM_REQUESTS});
do
    echo -e "\n${i}:\n"
    curl -v \
        -H "Authorization: ${QUOTAS_USER}" \
        -H "X-Game: ${QUOTAS_GAME}" \
        "${KRAKEND_ADDRESS}${QUOTAS_ENDPOINT}"
done
