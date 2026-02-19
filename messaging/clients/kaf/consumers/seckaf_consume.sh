#/bin/bash
export KTOPIC=portfolioupdates
# export KTOPIC=orderplacement
kaf consume \
    ${KTOPIC} \
    --config ./kaf_ssl_conf.yaml \
    --cluster seckafkabroker
