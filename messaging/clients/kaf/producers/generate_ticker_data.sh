#/bin/bash

# this script produces fake ticker prices, using the `prices.txt` data
# with each line for each of the list of tickers
KTOPIC=stockprice
KADDRESS=shrimp.ln:9092

echo -e "Producing: $KKEY into '$KTOPIC' topic\n" 

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
        kaf produce \
             ${KTOPIC} \
             -H Content-Type:application/json \
             -k ${KKEY} \
             -b ${KADDRESS} \
             -v \
             --input-mode full
    sleep $SLEEP_PERIOD
done < './prices.txt'
