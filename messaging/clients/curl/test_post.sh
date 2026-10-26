#!/bin/bash

export REPEAT=100

echo "Posting $REPEAT messages" 

for i in $(seq $REPEAT)
do 
    echo "Msg $i ----"
    . ./post_portfolio_order.sh
    SLEEP_TIME=$(expr $RANDOM % 3)
    sleep $SLEEP_TIME
    echo ""
done
