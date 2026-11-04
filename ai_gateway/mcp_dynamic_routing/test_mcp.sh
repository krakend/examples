#!/bin/bash

# Calls the same MCP tool once per planet. The only difference between
# the requests is the Host header, which KrakenD uses to pick the backend.

KRAKEND_URL=http://localhost:8080/fantasy

for PLANET in vegeta namek; do
  echo "--- Host: ${PLANET}"
  curl -s \
    -X POST "${KRAKEND_URL}" \
    -H "Host: ${PLANET}" \
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
  echo ""
done
