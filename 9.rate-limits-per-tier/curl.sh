#!/bin/bash

export TEST_URL='http://localhost:8080/test' 
echo -e "\nTesting GOLD tier:\n"
export HEADER_AUTHORIZATION="Authorization: Bearer gold-4d2c61e1-34c4-e96c-9456-15bd983c50"

curl \
    -H "${HEADER_AUTHORIZATION}" \
    $TEST_URL

