#!/bin/bash

export BASE_URL="localhost:7979"


req_paths=(
  "/dashboard"
  "/dashboard/"
  "/foo"
)

for i in {1..100}
do
    for rp in ${req_paths[@]}; do
        echo -e "\n"
        curl "${BASE_URL}${rp}"
    done
    sleep 1
done
