#/bin/bash
export KTOPIC=portfolioupdates
kaf consume \
    ${KTOPIC} \
    --config ./kaf_ssl_conf.yaml \
    --cluster seckafkabroker
