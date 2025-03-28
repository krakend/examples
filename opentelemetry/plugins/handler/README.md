# KrakenD - HTTP Handler Plugin Example

## Overview

The handler plugin, named `otel-handler`, demonstrates how to create a custom HTTP
handler that uses opentelemetry to:

- add attributes to current trace span: `plugin_value=<option_value>`
- create a new span named `plugin-otel-handler`
- and report a fake metric `handler_plugin.<option_value>|'default'.duration`

## File Structure

- `handler.go`: The main plugin code that implements the `RegisterHandler` interface.
- `go.mod`: The Go module file.
- `Makefile`: Instructions for building the plugin using Docker.

## Building the Plugin

You should build the plugin using the KrakenD Docker builder image. 

The `Makefile` includes targets for different architectures, to be used within
a docker environment (for example with `docker compose`).

### Steps

2. **Build for Different Architectures**:
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
               "url_pattern": "/__debug/"
            }
         ]
      }
   ],
   "extra_config": {
      "plugin/http-server": {
         "name": ["otel-handler"],
         "otel-handler": {
            "someOption": "some-value"
         }
      }
   }
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
