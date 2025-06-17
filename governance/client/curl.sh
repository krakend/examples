#!/bin/bash

KRAKEND_ADDRESS='http://localhost:8080'

# Number of requests to make
NUM_REQUESTS=1

# Endpoint to call
# QUOTAS_ENDPOINT='/request_weapon_power_up'
QUOTAS_ENDPOINT='/request_bomb_reload'
# QUOTAS_ENDPOINT='/consume_credit'

# Player to use
# QUOTAS_USER=bart
QUOTAS_USER=homer

# Select the game to use;
QUOTAS_GAME_SHIP=thunder
# QUOTAS_GAME_SHIP=spirit
# QUOTAS_GAME_SHIP=lighting

echo "Making ${NUM_REQUESTS} request(s), showing headers..."

for i in $(seq 1 ${NUM_REQUESTS});
do
    echo -e "\n${i}:\n"
    curl -v \
        -H "Authorization: ${QUOTAS_USER}" \
        -H "X-Game: ${QUOTAS_GAME_SHIP}" \
        "${KRAKEND_ADDRESS}${QUOTAS_ENDPOINT}"
done
