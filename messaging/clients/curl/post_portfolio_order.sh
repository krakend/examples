#!/bin/bash

export IKEY=$(date +%Y%m%d_%H%M%S)
export ACTION="buy"
export AMOUNT=$(expr $RANDOM % 100) 
export STOCK="AAPL"

echo "Publishing order: $ACTION $STOCK (amount: $AMOUNT)"

echo '{"action": "$ACTION", "amount": "$AMOUNT", "stock": "$STOCK"}' | \
    envsubst | \
    curl -X POST \
        -H "X-Idempotency-Key: $IKEY" \
        -H 'Some-Meta: avocados' \
        --data-binary @- \
        http://localhost:8080/portfolio/order
