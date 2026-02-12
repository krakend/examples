#/bin/bash

export KTOPIC=stockprices

kaf consume \
    ${KTOPIC} \
    -v \
    -b localhost:9092
