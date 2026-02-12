# Messaging

## Environment

To launch the test environment, just execute:

```
docker compose up -d 
```

It will bring up the following services:

- `kafkabroker`: a single node kafka broker
- `seckafkabroker`: a single node 
  
Along with services for telemetry:

- `grafana` 
- `tempo`
- `prometheus`

[##](##) Producing messages

In order to produce messages you will need to have installed the 
`kaf` tool.

# TODO: explain how to install the kaf tool

## "Story Telling"

The `kafkabroker` is a source of information for different market 
prices updates, so, we create a couple of topics where we will
publish some data:

- `stockprice`: some stock we are tracking

The `seckafkabroker` is our secure connection to the bank to place
market orders, and get updates of the status of our portfolio, so,
we secure it with a mTLS connection: a certificate is required to
connect to it.

- `orderplacement`
- `portfolioupdates`

