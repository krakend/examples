# Kafka Messaging Example

This example shows how KrakenD Enterprise interacts with Kafka in three ways:

- **Async agents** that consume a topic and send every message to a backend.
- **Publisher backends** that turn an HTTP request into a Kafka message.
- **Subscriber backends** that return the next message of a topic as the
  response of an HTTP request.

It uses two brokers: a plain one, and a secure one that requires mutual TLS
(mTLS) from its clients.

## The use case

Imagine you automate buying and selling stocks, based on custom logic that
lives in your backends:

- A public and reliable source publishes stock price updates in the
  `stockprice` topic of the `kafkabroker` broker.
- Your broker service lets you place buy and sell orders. To identify
  yourself, it gives you a client certificate that you use to connect with
  mTLS to `seckafkabroker`. You place orders in the `orderplacement` topic, and
  receive status updates about them in the `portfolioupdates` topic.

KrakenD connects to both brokers:

| Component                          | Broker           | Topic              | What it does                                                                 |
|------------------------------------|------------------|--------------------|------------------------------------------------------------------------------|
| `stocksconsumer` async agent       | `kafkabroker`    | `stockprice`       | Sends each price update to a backend                                         |
| `portfolioupdatesconsumer` async agent | `seckafkabroker` (mTLS) | `portfolioupdates` | Sends each order status update to a backend                       |
| `POST /ticker/publish`             | `kafkabroker`    | `stockprice`       | Publishes a price update (fake data source for the example)                  |
| `POST /portfolio/order`            | `seckafkabroker` (mTLS) | `orderplacement` | Places an order                                                        |
| `GET /portfolio/order`             | `seckafkabroker` (mTLS) | `orderplacement` | Returns the next pending order                                         |

The async agents send the messages to KrakenD's own `/__debug/` endpoint, so
you can see every message they process in the KrakenD logs.

## Requirements

- Docker and Docker Compose
- A KrakenD Enterprise license saved as `LICENSE` in this directory (need a
  trial license? [Contact us](https://www.krakend.io/contact-sales/))
- Optionally, the [`kaf`](https://github.com/birdayz/kaf) CLI, to produce
  and consume messages directly in Kafka

## Run the example

Generate the CA and the certificates. You only need to do this once (see
[`config/README.md`](./config/README.md) for details):

```bash
make setup
```

Start the environment:

```bash
make up
```

This starts the brokers, creates the topics with 3 partitions each, and then
starts the rest of the services:

| Service          | Address                                          | Purpose                                         |
|------------------|--------------------------------------------------|-------------------------------------------------|
| `krakend_ee`     | [http://localhost:8080](http://localhost:8080)   | KrakenD Enterprise                              |
| `kafkabroker`    | `localhost:9092`                                 | Plain Kafka broker                              |
| `seckafkabroker` | `localhost:49092`                                | Kafka broker that requires mTLS                 |
| `jaeger`         | [http://localhost:16686](http://localhost:16686) | Traces                                          |
| `grafana`        | [http://localhost:4000](http://localhost:4000)   | Dashboards (user `krakend`, password `krakend`) |
| `prometheus`     | [http://localhost:9090](http://localhost:9090)   | Metrics storage                                 |

Follow the KrakenD logs in a separate terminal to see the messages that the
async agents receive:

```bash
docker logs -f messaging-krakend
```

When you are done, stop everything with `make down`. To also delete the
generated certificates, run `make cleanup`.

## Try it

### Stock prices

Publish a fake price update every second through KrakenD:

```bash
bash clients/curl/generate_ticker_data.sh
```

The `stocksconsumer` async agent receives each update, and the KrakenD logs
show it as a request to `/__debug/stocksconsumer/received` with a body like
`{"ticker": "NVDA", "price": 134.41}`.

### Orders

Place an order:

```bash
bash clients/curl/post_portfolio_order.sh
```

Read the next pending order:

```bash
bash clients/curl/get_portfolio_order.sh
```

The response contains the order in the body, and the message metadata
(`X-Idempotency-Key`, `Some-Meta`, ...) as headers. When there are no pending
orders, the request fails after a short timeout.

[`test_post.sh`](./clients/curl/test_post.sh) and
[`test_get.sh`](./clients/curl/test_get.sh) repeat these calls 100 times, so
you can run them in two terminals at the same time.

### Portfolio updates and direct Kafka access

The scripts in [`clients/kaf`](./clients/kaf) use `kaf` to talk to the brokers
directly, without KrakenD:

- [`producers/send_portfolio_update.sh`](./clients/kaf/producers/send_portfolio_update.sh):
  publishes a message in `portfolioupdates`, which the
  `portfolioupdatesconsumer` async agent then receives.
- [`producers/generate_ticker_data.sh`](./clients/kaf/producers/generate_ticker_data.sh)
  and [`producers/kaf.sh`](./clients/kaf/producers/kaf.sh): publish price
  updates in `stockprice`.
- [`consumers/kaf_consume.sh`](./clients/kaf/consumers/kaf_consume.sh): reads
  the `stockprice` topic.
- [`consumers/seckaf_consume.sh`](./clients/kaf/consumers/seckaf_consume.sh):
  reads a topic of the secure broker, `orderplacement` by default
  (`./seckaf_consume.sh portfolioupdates` to read the other one).

If you don't have `kaf`, you can publish a portfolio update with the Kafka
tools that come with the broker image:

```bash
echo '{"order":"k1","status":"executed"}' | docker exec -i messaging-seckafkabroker \
    /opt/kafka/bin/kafka-console-producer.sh --bootstrap-server localhost:9092 --topic portfolioupdates
```

## Observability

KrakenD reports traces to Jaeger under the `krakend_messaging` service, and
metrics to Prometheus. Besides the usual KrakenD metrics, you get the
`messaging_read_*` and `messaging_write_*` metrics with the size and duration
of the messages read and written, per `kind` and `topic`.

## Learn more

- [Kafka driver for the Async Agent](https://www.krakend.io/docs/enterprise/async/kafka/),
  including how KrakenD acknowledges Kafka messages
- [Kafka PubSub publisher and subscriber](https://www.krakend.io/docs/enterprise/backends/pubsub/kafka/),
  including the producer options, partition selection, and the OpenTelemetry
  metrics and traces
