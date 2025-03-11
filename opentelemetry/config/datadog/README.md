# Using OTEL with DataDog

Based on the [OpenTelemetry in DataDog](https://docs.datadoghq.com/opentelemetry/) 
documentation there are two ways of sending the OTEL metrics and traces to 
datadog. 

With this example we provide examples of both ways to do it.

## Collector

Under the `collector` directory you can find the [instructions](./collector/README.md)
and example files to send OTEL metrics and traces.

## Agent

Unde the `agent` directory you can find the [instructions](./agent/README.md)
and example files to use the DataDog agent to send the OTEL
metrics and traces.
