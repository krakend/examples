#!/bin/bash

# this script produces fake ticker prices, using the `prices.txt` data
# with each line for each of the list of tickers
KRAKEND_ADDRESS=localhost:8080
SLEEP_PERIOD=1
TICKERS=('NVDA', 'AAPL', 'AMZN')

export idx=0
while read -r line
do
    export idx=$(expr $idx + 1)
    export idx=$(expr $idx % 3)

    export PRICE=$line
    export TICKER=${TICKERS[$idx]}
    export KKEY=$(date +%Y%m%d_%H%M%S)
    echo '{"ticker": "$TICKER", "price": $PRICE}' | envsubst | \
        curl -X POST \
             -H 'Content-Type:application/json' \
             -H "X-Ticker: $TICKER" \
             -H "X-Time: $KKEY" \
             --data-binary @- \
             http://${KRAKEND_ADDRESS}/ticker/publish
    sleep $SLEEP_PERIOD
done < './prices.txt'
