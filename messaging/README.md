# Messaging

## Environment

The first you running, you will need to generate the CA and the
self signed certificates. Run `make setup` and fill the required fields
and accept signing the  certificates. (see [config's README.md](./confit/README.md)).

It will bring up the following services:

- `kafkabroker`: a single node kafka broker
- `seckafkabroker`: a single node kafka broker configured to use
    mTLS for its connections

Along with services for telemetry:

- `prometheus`
- `grafana`
- `jaeger`

## The example "use case"

Imagine we want to automate some stocks buying and selling, based on
some custom logic that is implemented in some of our backends. 

We have a source of public reliable information that publishes stock
prices updates using a kafka service: that is `kafkabroker`. This service
has a `stockprice` topic, where we can subscribe to receive the updates.

We also have a service that allows to manage our portfolio and execute
buy an sell order. To identify ourselves we are given a client certificate
that we will use to connect using `mTLS` to the service. That is 
`seckafkabroker`. This service exposes a topic where we can place our
orders: `orderplacement`, and also offers us a topic `portfolioupdates` so we 
get status updates on the executed orders.


### Producing messages

In order to produce messages you will need to have installed the
`kaf` tool: [https://github.com/birdayz/kaf](https://github.com/birdayz/kaf).


Under the [`clients/kak/producer`](./clients/kaf/producer) folder you will find
scripts to generate data:

- [`generate_ticker_data.sh`](./clients/kaf/producers/generate_ticker_data.sh): will
    emit messages for ticker from the `prices.txt` file every 1 second.
    (we can change the period by changing the `SLEEP_PERIOD` variable).
- [`send_portfolio_update.sh`](./clients/producer/kaf_ssl.sh): to emit a fake
    message in the `portfolioupdates` topic.
- [`kaf.sh`](./clients/producer/kaf.sh): to emit a single ticker value
    to `stockprice` topic on the `kafkabroker` server.

### Consuming messages

To chck the messages that are in the queues there are some scripts
in the `clients/kaf/consumers` directory:

- `kaf_consume.sh`: to see the `stockprice` topic 
- `seckaf_consume.sh`: to see what is in the seckafkabroker topics
    (edit the script to select the topic to read from).


# Improved Kafka support (Documentation)

## Async Agent

The `async_agent` section in the main config contains an array of objects that
define the configuration of a given async agent. By using the `async/kafka`
namespace under the `extra_config` key, we select the Kafka driver.

The driver fields are the same that can be found in the kafka pubsub 
subscriber config `reader` field:

- `group_id`: the name of the consumer group to use
- `key_meta`: the name of the header where the kafka message key value is written
- `topics`: a list of topics to read from.
- `connection`: details to connect to the kafka service
- `consumer`: details about to read from the topic


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

- `group_session_timeout`: The timeout used to detect consumer failures when using Kafka's group management facility.
	The consumer sends periodic heartbeats to indicate its liveness to the broker.
	If no heartbeats are received by the broker before the expiration of this session timeout,
	then the broker will remove this consumer from the group and initiate a rebalance.
	Note that the value must be in the allowable range as configured in the broker configuration
	by `group.min.session.timeout.ms` and `group.max.session.timeout.ms` (default 10s)
    
    
- `group_heartbeat_interval`: The expected time between heartbeats to the consumer coordinator when using Kafka's group
	management facilities. Heartbeats are used to ensure that the consumer's session stays active and
	to facilitate rebalancing when new consumers join or leave the group.
	The value must be set lower than Consumer.Group.Session.Timeout, but typically should be set no
	higher than 1/3 of that value.
	It can be adjusted even lower to control the expected time for normal rebalances (default 3s)


- `group_rebalance_strategies`: the priority-ordered list of client-side consumer group
	balancing strategies that will be offered to the coordinator. The first
	strategy that all group members support will be chosen by the leader.
	default: [ NewBalanceStrategyRange() ]
	can be :
	range -> RangeBalanceStrategyName
	roundrobin -> RoundRobinBalanceStrategyName
	sticky -> StickyBalanceStrategyName
- `group_rebalance_timeout`: The maximum allowed time for each worker to join the 
    group once a rebalance has begun.
	This is basically a limit on the amount of time needed for all tasks to flush any pending
	data and commit offsets. If the timeout is exceeded, then the worker will be removed from
	the group, which will cause offset commit failures (default 60s).
- `group_instance_id`: support KIP-345

- `fetch_default`: The default number of message bytes to fetch from the broker in each
	request (default 1MB). This should be larger than the majority of
	your messages, or else the consumer will spend a lot of time
	negotiating sizes and not actually consuming. Similar to the JVM's
	`fetch.message.max.bytes`.

- `isolation_level`: IsolationLevel support 2 mode:
	- use `read_commited` (default) to consume and return all messages in message channel
	- use `read_uncommited` to hide messages that are part of an aborted transaction

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
	- `"no_response"`: no requied acks (same as `"0"`)
	- `"wait_for_local"`: waits for only the local commit to succeed before responding.
        (same as `"1"`).
	- `"wait_for_all"`: waits for all in-sync replicas to commit before responding.
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


### Consumer config

It only contains a  `reader` field, that corresponds to the same configuration
that the async kafka driver uses.

- `reader`:  same configuration that the async kafka driver uses.

## OpenTelemery

If OTEL is enabled these metrics will be reported by default for all messages
read (either in the async agent, or a subscriber backend) and written (by
a publisher backend).

All of these matrics have these attributes set:
    - `kind`: that helps identify the kind of queue systemthe message is
        using (for example `kafka`).
    - `topic`: the topic where the message is read from or going to be 
        writen to.
        
- For reading: 
	- `messaging.read.body.size`: histogram of body sizes in bytes (does not include the size of metadata)
	- `messaging.read.body.duration`: histogram of the duration taken to read a message
	- `messaging.read.ack.duration`
	- `messaging.read.failure.count`
  
- For writing:
    - `messaging.write.body.size`: histogram of body sizes in bytes (does not include the size of metadata)
    - `messaging.write.body.duration`:  histogram of duration taken to write a message
    - `messaging.write.failure.count`: count of messages failed to be written


## **ACK** behaviour in Kafka

TODO: put here the explanation of not acking a message, bug acking the one that comes after it
