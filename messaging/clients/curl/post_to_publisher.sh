#!/bin/bash

curl -X POST \
    -H 'X-Idempotency-Key: 0000001' \
    -H 'Some-Meta: avocados' \
    -d '{"foo": "bar"}' \
    http://localhost:8080/portfolio/order
