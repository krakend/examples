# OpenTelemetry example

This example shows how use OpenTelemetry to report at different 
rates to two different exporters:

- a local Jaeger instance (will be availabe at: http://localhost:16686/search ).
- to DataDog, using your own API key (optional)

You need a KrakenD Enterprise license saved as `config/krakend/LICENSE`.

The example also uses three custom plugins that report their own spans. Build
them first with the `make amd64` (or `make arm64`) target of each plugin folder
(see the [plugins README](./plugins/README.md)); the resulting `.so` files are
placed in `config/krakend/plugins`.

Then start the environment:

```
bash ./run.sh
```

To also send the telemetry to DataDog, set your API key in the `DD_API_KEY`
env var ( `export DD_API_KEY=<your_api_key>` ) before running the script, so a
configuration file is generated for the OpenTelemetry collector. Without the
key, the collector receives the telemetry but does not export it, and you can
still see the traces in Jaeger.

Once it is running you can send some requests using the script in
`./client/curl_requests.sh`, to get some data to look at.


## Plugins

To check how to use OpenTelemetry in [plugin's README](./plugins/README.md)

## DataDog

To setup reporting to DataDog check the [DataDog README](./config/datadog/README.md).
