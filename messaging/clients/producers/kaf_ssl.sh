#/bin/bash

kaf produce \
    krafka \
    --config ./kaf_ssl_conf.yaml \
    --cluster seckafkabroker \
    -H Content-Type:application/json \
    -k $(date +%Y%m%d_%H%M%S) \
    -v \
    --input-mode full < kaf_payload.json
