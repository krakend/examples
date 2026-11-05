#!/bin/bash

cd "$(dirname "$0")" || exit 1
export REPEAT=100

echo "Getting $REPEAT messages" 

for i in $(seq $REPEAT)
do 
    echo "Msg $i ----"
    . ./get_portfolio_order.sh
    sleep 1
    echo ""
done
