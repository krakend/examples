#!/bin/bash

export HOST=localhost
export HOST=vegeta
export HOST=namek

curl -i \
  -X POST http://${HOST}:8088/fantasy \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "method": "tools/call",
    "params": {
      "name": "get_fantasy_planet",
      "arguments": {
        "landscape": "mountains"
      }
    },
    "id": 1
  }'
