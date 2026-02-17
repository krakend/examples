# Messaging

## Environment

First of all, enter the `config` dir, and follow the 
[README.md](./config/README.md) instructions to generate a self
signed certificate for the client and the `seckafkabroker` service,
to be able to use `mTLS`

Once done that, launchthe  test  environment with:

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
`kaf` tool: [https://github.com/birdayz/kaf](https://github.com/birdayz/kaf).

Under the [`clients/producer`](./clients/producer) folder you will find
two scripts:

- [`kaf.sh`](./clients/producer/kaf.sh): to generate fake data for the 
    `stockprice` topic on the `kafkabroker` server.
- [`kaf_ssh.sh`](./clients/producer/kaf_ssl.sh): to generate fake data for the 
    `portfolioupdates` topic on the `seckafkaborker`.
    

### Kaf to produce message with SSL TLS

https://github.com/birdayz/kaf/blob/master/examples/ssl_keys.yaml


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

