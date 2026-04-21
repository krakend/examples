#!/bin/sh

docker compose up -d 
echo "waiting for startup ..."
echo "5"
sleep 1
echo "4"
sleep 1
echo "3"
sleep 1
echo "2"
sleep 1
echo "1"
sleep 1

echo "creating topics:"

echo "creating 'stockprice' topic"
docker exec messaging-kafkabroker \
    /opt/kafka/bin/kafka-topics.sh \
    --bootstrap-server localhost:9092 \
    --create \
    --topic stockprice \
    --partitions 3 

echo "creating 'orderplacement' topic\n"
docker exec messaging-seckafkabroker \
    /opt/kafka/bin/kafka-topics.sh \
    --bootstrap-server localhost:9092 \
    --create \
    --topic orderplacement \
    --partitions 3 

echo "creating 'portfolioupdates' topic\n"
docker exec messaging-seckafkabroker \
    /opt/kafka/bin/kafka-topics.sh \
    --bootstrap-server localhost:9092 \
    --create \
    --topic portfolioupdates \
    --partitions 3 
