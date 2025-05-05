#!/bin/bash

# TODO: down here the path to your "hey" executable:
# https://github.com/rakyll/hey
export HEY_BIN="./hey_linux_amd64"
export TEST_URL='http://localhost:8080/test' 

echo -e "\nTesting GOLD tier:\n"
export HEADER_AUTHORIZATION="Authorization: Bearer gold-4d2c61e1-34c4-e96c-9456-15bd983c50"
$HEY_BIN -n 20 -c 10 \
    -H "${HEADER_AUTHORIZATION}" \
    $TEST_URL


echo -e "\n\nTesting SILVER tier:\n"
export HEADER_AUTHORIZATION="Authorization: Bearer silver-4d2c61e1-34c4-e96c-9456-15bd983c50"
$HEY_BIN -n 20 -c 10 \
    -H "${HEADER_AUTHORIZATION}" \
    $TEST_URL
