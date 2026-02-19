#!/bin/bash

export IKEY=$(date +%Y%m%d_%H%M%S)
curl -X POST \
    -H "X-Idempotency-Key: $IKEY" \
    -H 'Some-Meta: avocados' \
    -d '{"foo": "bar"}' \
    http://localhost:8080/portfolio/order
