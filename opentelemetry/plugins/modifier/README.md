# KrakenD - Request/Response Modifier Plugin Example

This repository provides an example of how to build and inject a Request/Response Modifier Plugin into the KrakenD API Gateway. Follow the steps below to understand how to set up and use the modifier plugin.

## Overview

The modifier plugin, named `otel-modifier`, demonstrates how to create custom request and response modifiers that uses opentelemetry to: 

- add attributes to current trace span: `plugin_req_mod_option=<option_value>` and `plungin_resp_mod_option`.
- create a new span named `req_mod_plugin_span` and `resp_mod_plugin_span`.

Usage of metrics can work the same way as in the [handler plugin](../handler/README.md)

## File Structure

- `modifier.go`: The main plugin code that implements the request and response modifiers.
- `go.mod`: The Go module file.
- `Makefile`: Instructions for building the plugin using Docker.


## Building the Plugin

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
         "method": "GET",
         "extra_config": {
            "plugin/req-resp-modifier": {
               "name": ["otel-modifier-request", "otel-modifier-response"],
               "otel-modifier-request": {
                  "option": "value"
               },
               "otel-modifier-response": {
                  "option": "value"
               }
            }
         },
         "backend": [
            {
               "url_pattern": "/__debug/"
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


For more details on request/response modifier plugins and extending KrakenD, refer to the [official documentation](https://www.krakend.io/docs/extending/plugin-modifiers/).
