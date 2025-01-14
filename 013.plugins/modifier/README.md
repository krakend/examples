# KrakenD - Request/Response Modifier Plugin Example

This repository provides an example of how to build and inject a Request/Response Modifier Plugin into the KrakenD API Gateway. Follow the steps below to understand how to set up and use the modifier plugin.

## Overview

The modifier plugin, named `my-modifier`, demonstrates how to create custom request and response modifiers that can be integrated into KrakenD. The plugin allows you to intercept and modify the request before it reaches the backend and the response before it is returned to the client.

## File Structure

- `modifier.go`: The main plugin code that implements the request and response modifiers.
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
    cd examples/plugins/modifier
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
         "method": "GET",
         "extra_config": {
            "plugin/req-resp-modifier": {
               "name": ["my-modifier-request", "my-modifier-response"],
               "my-modifier-request": {
                  "option": "value"
               },
               "my-modifier-response": {
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

## Custom Logic

The `my-modifier` plugin adds custom logic by intercepting and potentially modifying the request and response. The example plugin logs details about the request and response but does not modify them.

## Usage

1. **Build the plugin** using the appropriate make target. The plugin will be generated under the `plugins` folder.
2. **Test your KrakenD using Docker**, for instance:
```bash
docker run --rm -it --name krakend -p 8080:8080 -v "$PWD:/etc/krakend" devopsfaith/krakend
```
3. **Send an example call**
```bash
curl -iG 'http://localhost:8080/example'
```

For more details on request/response modifier plugins and extending KrakenD, refer to the [official documentation](https://www.krakend.io/docs/extending/plugin-modifiers/).
