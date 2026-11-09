# KrakenD - HTTP Handler Plugin Example

> [!WARNING]
> **Deprecated: Go plugins are no longer supported in the Community Edition.**
> KrakenD CE 3.0 [dropped custom Go plugins](https://www.krakend.io/blog/dropping-plugins-support-on-community/), so this example only works with KrakenD CE 2.x, and is pinned to `krakend:2.13.11` and `krakend/builder:2.13.11`, the last release supporting them.
> To keep using Go plugins on 3.0, move to [KrakenD Enterprise](https://www.krakend.io/docs/enterprise/extending/), which also changes the plugin configuration format (see the [`opentelemetry`](../../opentelemetry/plugins) example for working 3.0 plugins). For custom logic on the Community Edition, use [Lua scripting](https://www.krakend.io/docs/endpoints/lua/).

This repository provides an example of how to build and inject a Handler Plugin into the KrakenD API Gateway. Follow the steps below to understand how to set up and use the handler plugin.

## Overview

The handler plugin, named `my-handler-plugin`, demonstrates how to create a custom HTTP handler that can be integrated into KrakenD. The plugin replaces the default HTTP handler by adding custom logic before and after the request is processed.

## File Structure

- `handler.go`: The main plugin code that implements the `RegisterHandler` interface.
- `go.mod`: The Go module file.
- `Makefile`: Instructions for building the plugin using Docker.

## Prerequisites

- [Go](https://golang.org/dl/)
- Docker
- KrakenD API Gateway (version 2.x)

## Building the Plugin

You should build the plugin using the KrakenD Docker builder image. The `Makefile` includes targets for different architectures.

### Steps

1. **Clone the Repository**:

    ```bash
    git clone https://github.com/krakend/examples.git
    cd examples/plugins/handler
    ```

2. **Build for Different Architectures**:
    - For **amd64**:

        ```bash
        make amd64
        ```

    - For **arm64**:

        ```bash
        make arm64
        ```

    - For **linux_amd64** (non-docker):

        ```bash
        make linux_amd64
        ```

    - For **linux_arm64** (non-docker):

        ```bash
        make linux_arm64
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
         "name": ["my-handler-plugin"],
         "my-handler-plugin": {
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

## Custom Logic

The `my-handler-plugin` adds custom logic by wrapping the default HTTP handler. It logs the configuration option and then processes the request using the default handler.

## Usage

1. **Build the plugin** using the appropriate make target. The plugin will be generated under the `plugins` folder.
2. **Test your KrakenD using Docker**, for instance:
```bash
docker run --rm -it --name krakend -p 8080:8080 -v "$PWD:/etc/krakend" krakend:2.13.11
```
3. **Send an example call**
```bash
curl -iG 'http://localhost:8080/example'
```

For more details on handler plugins and extending KrakenD, refer to the [official documentation](https://www.krakend.io/docs/extending/http-server-plugins/).
