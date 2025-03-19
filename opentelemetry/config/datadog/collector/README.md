# Using OTEL Collector to send metrics and traces to DataDog

In order to generate a config with your DataDog API Key, you
must setup a `DD_API_KEY` and have `envsubst` installed, and
just run `./run.sh`


- [Extensive example of collector config for datadog](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/exporter/datadogexporter/examples/collector.yaml)

References:

https://www.datadoghq.com/blog/datadog-agent-with-otel-collector/

https://docs.datadoghq.com/opentelemetry/

https://docs.datadoghq.com/opentelemetry/interoperability/otlp_ingest_in_the_agent/?tab=host

https://docs.datadoghq.com/opentelemetry/collector_exporter/configuration/

https://opentelemetry.io/docs/collector/configuration/

https://docs.datadoghq.com/opentelemetry/collector_exporter/deployment/?tab=localhost#docker

https://docs.datadoghq.com/opentelemetry/collector_exporter/deployment/?tab=localhost
