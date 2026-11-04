#!/bin/bash

export BASE_URL="localhost:8080"

# req_paths=(
#   "/dashboard"
#   "/dashboard/"
#   "/foo"
# )
req_paths=(
    "/dashboard?foo=3&bar=10"
    "/dashboard?mytext=something"
    "/dashboard?a=b"
    "/otelplugins"
    "/otelplugins?a=b"
)

for i in {1..200}
do
    for rp in ${req_paths[@]}; do
        echo -e "\n"
        curl "${BASE_URL}${rp}"
    done
    sleep 1
done
