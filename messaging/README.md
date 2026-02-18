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

# Improved Kafka support

## Async Agent

The `async_agent` section in the main config contains an array of objects that
define the configuration of a given async agent. By using the `async/kafka`
namespace under the `extra_config` key, we select the Kafka driver.

- `group_id`: the name of the consumer group to use
- `topics`: a list of topics to read from 
- `connection`: details 


#### The `connection` config

- `brokers`: (required) the list of `[host]:[port]` addresses of the kafka brokers to
    connecto to.
- `client_id`: (optional) a name to identify the client that is stablishing the
    connection  (nothing to do with the consumer group id). If not provided,
    it defaults to Krakend with the version.
- `client_tls`: (optional) in order to authenticate using `mTLS`, or allow
    https insecure connection, we have the same options available that
    in regular [https backend requests](https://www.krakend.io/docs/enterprise/service-settings/tls/#client-tls-settings).
- `sasl`: (optional) to authenticate using [Simple Authentication and Security Layer](https://en.wikipedia.org/wiki/Simple_Authentication_and_Security_Layer).
	There are multiple SASL authentication methods but the current implementation
    is **limited to plaintext (SASL/PLAIN) authentication**.
    
    - `user`: the username to use
    - `password`: the password
    - `azure_event_hub`: (optional) set to true if you want to use Azure
        EventHub, that uses SASL V0 (instead of the default V1.x).

- `dial_timeout`: (optional) specifies a duration (in string format: "100ms", "1s",..)
	for the maximum time to dial to the brokers

- `read_timeout`: (optional) specifies a duration (in string format: "100ms", "1s",..)
	for the maximum time complete a read operation.


- `write_timeout`: (optional) specifies a duration (in string format: "100ms", "1s",..)
	for the maximum time complete a write operation.

- `keep_alive`: (optional) specifies a duration (in string format: "100ms", "1s",..)
	to maintain a connection open while while not being used.

- `rack_id`: (optional) a rack identifier for this client. This can be any string 
    value which indicates where this client is physically located. It 
    corresponds with the broker config 'broker.rack'

- `channels_buffer_size`: (optional) The number of events to buffer in internal 
    and external channels. This permits the producer and consumer to continue 
    processing some messages in the background while user code is working,
    greatly improving throughput. Defaults to 256.

### The `consumer` config

Since the `topic` to consume from is already defines at the general async
agent config level, here we define properties specific for the kafka driver.



## Pub Sub

### Publisher

By adding the `backend/pubsub/publisher/kafka` key under a **backend**'s
`extra_config`, we override the regular configuration of the backend
to publish a message in a kafka queue.

The `input_headers` of the backend will be converted to the message
metadata.

The configuration has the following properties:

- `success_status_code`: the return status code after a message has been
    successfully published into a topic. **Warning**: is only possible
    to specify a value of `2xx`, however, KrakenD only identifies a
    successful response the one that return `200` or  `201`.

- `writer`: the definition of the topic we want to write to
    - `connection`: this structure is the same that in
    - `producer`: contains the specific configuration related to how to 
        configure 
    - `topic`: the topic name to write to
    - `key_meta`: the meta field that will be used as key of the message

#### `producer` object config

- `max_message_bytes`: (int > 0, optional) The maximum permitted size of a message 
    (defaults to 1000000). Should be set equal to or smaller than 
    the broker's `message.max.bytes`.

- `required_acks`: (string, optional) The level of acknowledgement reliability 
    needed from the broker (defaults to WaitForLocal). Equivalent to 
    the `request.required.acks` setting of the  JVM producer.
	- "no_response" -> NoResponse RequiredAcks = 0
	- "wait_for_local" -> waits for only the local commit to succeed before responding.
	- "wait_for_all" -> waits for all in-sync replicas to commit before responding.
	The minimum number of in-sync replicas is configured on the broker via
	the `min.insync.replicas` configuration key. The string can also contain
    a number > 0 to set a given number of expected acks.

- `required_acks_timeout`: (duration, optional) The maximum duration 
    the broker will wait the receipt of the number of `required_acks`
	(defaults to 10 seconds). This is only relevant when `required_acks`
	is set to WaitForAll or a number > 1. Only supports
	millisecond resolution, nanoseconds will be truncated. Equivalent to
	the JVM producer's `request.timeout.ms` setting.

- `compression_codec`: (string, optional) The type of compression to use on 
    messages (defaults to no compression).
	Similar to `compression.codec` setting of the JVM producer.
	[ `"none"`, `"gzip"`, `"snappy"`, `"lz4"`, `"zstd"` ] (defaults to `none`)
- `compression_level`: (int, optional) The level of compression to use on 
    messages. The meaning depends on the actual compression type used and 
    defaults to default compression level for the codec.
- `idempotent`: (boolean, optional) If enabled, the producer will ensure 
    that exactly one copy of each message is written.
- `retry_max`: (int, optional) The total number of times to retry sending a message (default 3).
	Similar to the `message.send.max.retries` setting of the JVM producer.
- `retry_backoff`: (duration, optional) How long to wait for the cluster to 
    settle between retries (default 100ms). Similar to the `retry.backoff.ms` 
    setting of the JVM producer.


### Consumer
