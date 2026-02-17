#/bin/bash
export KTOPIC=stockprice
kaf consume \
    ${KTOPIC} \
    -v \
    -b localhost:9092
