#!/bin/bash

echo "------------------------------------------"
echo " Show Fake API successful payload:"
echo "------------------------------------------"
curl -i http://localhost:8088/data/
echo ""
echo "******************************************"
echo ""

echo "------------------------------------------"
echo " Show Random Scorer results:"
echo "------------------------------------------"
for i in $(seq 4)
do 
    echo "  * req num ${i}"
    curl -i -X POST http://localhost:8088/scorer/
    echo "______________________________________"
done
echo "******************************************"
echo ""

