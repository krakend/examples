# KrakenD OpenTelemetry in Plugins Examples

This folder contains examples about how to use configured opentelemetry exporters
in your own custom plugins. 

These examples cover two types of plugins: Handler/Server, and Request/Response Modifier.

## Limitations

- You can use the globally configured exporters, but plugins **cannot use the exporters overrides** feature
    that are configured at the endpoint level. 
    
- Ih the handler / server traces, the span you create can be the parent for the
    rest of spans down the KrakenD pipeline. However, in Request/Response modifiers
    since the modified context is not propagated, if you create a new span it will
    be displayed as a sibling to the rest of the spans. In both kind of plugins you
    can set attributes to the existing span in the context.
    
## Overview

There are two main directories, each containing a specific type of plugin:

### 2. Handler/Server Plugin

The handler/server plugin shows how to create a custom HTTP handler, that sets
custom attributes in the incoming / existing trace span, and also creates a new 
span, before calling the rest of the pipeline.

**More details and instructions can be found in the [`handler` README](./handler/README.md).**

### 3. Request/Response Modifier Plugin

The request/response modifier plugin illustrates how to create custom request and response modifiers
that sets custom attributes in the incoming / exisint trace span. his plugin can modify the request before it is sent to the backend and the response before it is returned to the client.

**More details and instructions can be found in the [`modifier` README](./modifier/README.md) directory.**

## Documentation

For more information on extending KrakenD with plugins, refer to the official documentation:

- [HTTP Client Plugins](https://www.krakend.io/docs/extending/http-client-plugins/)
- [HTTP Server Plugins](https://www.krakend.io/docs/extending/http-server-plugins/)
- [Request/Response Modifier Plugins](https://www.krakend.io/docs/extending/plugin-modifiers/)

