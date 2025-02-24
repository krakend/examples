#!/bin/bash
if command -v jq 2>&1 > /dev/null
then
    curl localhost:8080/grpc/travel | jq
else 
    curl localhost:8080/grpc/travel
fi
