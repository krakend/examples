# KrakenD - HTTP Client Plugin Example

This directory provides an example of how to build and inject a Client Plugin into the KrakenD API Gateway,
that can use the existing telemetry to continue reporting traces and metrics.

## Overview

The client plugin, named `otel-client`, demonstrates how to create a custom HTTP client 
that can keep showing traces and reporting metrics. 

## File Structure

- `client.go`: The main plugin code that implements the `RegisterClient` interface.
- `go.mod`: The Go module file.
- `Makefile`: Instructions for building the plugin using Docker.


You should build the plugin using the KrakenD Docker builder image. The `Makefile` includes targets for different architectures.

### Steps

**Build for Different Architectures**:
    - For **amd64**:

        ```bash
        make amd64
        ```

    - For **arm64**:

        ```bash
        make arm64
        ```

## Plugin Configuration

To use the plugin in your KrakenD configuration, add it under the `extra_config` section of your `krakend.json` file.

### Example Configuration

```json
{
  "version": 3,
  "plugin": {
    "pattern": ".so",
    "folder": "/etc/krakend/plugins"
  },
  "host": ["http://localhost:8080/"],
  "debug_endpoint": true,
  "endpoints": [
    {
      "endpoint": "/example",
      "backend": [
        {
          "url_pattern": "/__debug/",
          "extra_config": {
            "plugin/http-client": {
              "name": "otel-client",
              "otel-client": {
                "option": "/some-path"
              }
            }
          }
        }
      ]
    }
  ]
}

```

## Logger Interface

The plugin supports a logger interface to help with debugging and logging messages.

### Logger Methods

- `Debug(v ...interface{})`
- `Info(v ...interface{})`
- `Warning(v ...interface{})`
- `Error(v ...interface{})`
- `Critical(v ...interface{})`
- `Fatal(v ...interface{})`


For more details on client plugins and extending KrakenD, refer to the [official documentation](https://www.krakend.io/docs/extending/http-client-plugins/).
