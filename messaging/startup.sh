#!/bin/sh

# Starts the Kafka brokers first and creates the topics with 3 partitions
# before KrakenD starts. Otherwise KrakenD would connect first and Kafka
# would auto-create the topics with a single partition.

KAFKA_TOPICS=/opt/kafka/bin/kafka-topics.sh

docker compose up -d kafkabroker seckafkabroker

wait_for_broker() {
    echo "waiting for $1 ..."
    until docker exec "$1" ${KAFKA_TOPICS} --bootstrap-server localhost:9092 --list > /dev/null 2>&1
    do
        sleep 1
    done
}

create_topic() {
    echo "creating '$2' topic in $1"
    docker exec "$1" ${KAFKA_TOPICS} \
        --bootstrap-server localhost:9092 \
        --create \
        --if-not-exists \
        --topic "$2" \
        --partitions 3
}

wait_for_broker messaging-kafkabroker
wait_for_broker messaging-seckafkabroker

create_topic messaging-kafkabroker stockprice
create_topic messaging-seckafkabroker orderplacement
create_topic messaging-seckafkabroker portfolioupdates

docker compose up -d
