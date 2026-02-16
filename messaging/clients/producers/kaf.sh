#/bin/bash

KTOPIC=stockprice
KADDRESS=shrimp.ln:9092

echo -e "Producing: $KKEY into '$KTOPIC' topic\n" 

SLEEP_PERIOD=5
PRICES=("309.28", "312.59", "315.7", "314.68", "311.9", "309.70", "319.00")

for p in ${PRICES[@]}; do
    echo  "producing price ${p}" 
    export PRICE=$p
    export KKEY=$(date +%Y%m%d_%H%M%S)
    envsubst < kaf_payload.json | \
        kaf produce \
            ${KTOPIC} \
            -H Content-Type:application/json \
            -k ${KKEY} \
            -b ${KADDRESS} \
            -v \
            --input-mode full
    sleep ${SLEEP_PERIOD}
done 

