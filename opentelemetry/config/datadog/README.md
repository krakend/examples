# Using OTEL with DataDog

Based on the [OpenTelemetry in DataDog](https://docs.datadoghq.com/opentelemetry/) 
documentation there are two ways of sending the OTEL metrics and traces to 
datadog. 

With this example we provide an example using the collector.

## Collector

Under the `collector` directory you can find the [instructions](./collector/README.md)
and example files to send OTEL metrics and traces.

## Agent

Under the `agent` directory you can find the [README](./agent/README.md)
with an example file that could be used to run the agent to collect
metrics. (However, we do not provide a working example to use it).
