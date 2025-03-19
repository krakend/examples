# OpenTelemetry example

This example shows hot use OpenTelemetry to report at different 
rates to two different exporters:

- a local Jaeger instance (will be availabe at: http://localhost:16686/search ).
- to DataDog, usin your own API key
  
You must set your api key to the `DD_API_KEY` env var ( `export DD_API_KEY=<your_api_key>` ), before running the script to start the docker compose,
so a file is generated for the datadog agent.

```
bash ./run.sh
```

Once it is running you can send some requests using the script in
`./client/curl_request.sh`, to get some data to look at.


## Plugins

To check how to use OpenTelemetry in [plugin's README](./plugins/README.md)

## DataDog

To setup reporting to DataDog check the [DataDog README](./config/datadog/README.md).
