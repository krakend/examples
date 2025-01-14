# KrakenD Plugin Examples

This repository contains examples demonstrating how to build and inject custom plugins into the KrakenD API Gateway. These examples cover three types of plugins: HTTP Client, Handler/Server, and Request/Response Modifier.

## Overview

The repository is structured into three main directories, each containing a specific type of plugin:

### 1. Client Plugin

The client plugin demonstrates how to create a custom HTTP client that can be integrated into KrakenD. The plugin adds custom logic for handling HTTP requests.

**More details and instructions can be found in the `client` directory.**

### 2. Handler/Server Plugin

The handler/server plugin shows how to create a custom HTTP handler. This plugin allows you to add custom logic to process HTTP requests before they reach the backend.

**More details and instructions can be found in the `handler` directory.**

### 3. Request/Response Modifier Plugin

The request/response modifier plugin illustrates how to create custom request and response modifiers. This plugin can modify the request before it is sent to the backend and the response before it is returned to the client.

**More details and instructions can be found in the `modifier` directory.**

## Building the Plugins

Each plugin directory contains the necessary code and a `Makefile` for building the plugin using Docker. Follow the instructions in each directory's README to build and integrate the plugins into your KrakenD setup.

## Documentation

For more information on extending KrakenD with plugins, refer to the official documentation:

- [HTTP Client Plugins](https://www.krakend.io/docs/extending/http-client-plugins/)
- [HTTP Server Plugins](https://www.krakend.io/docs/extending/http-server-plugins/)
- [Request/Response Modifier Plugins](https://www.krakend.io/docs/extending/plugin-modifiers/)

This repository aims to provide a starting point for creating your own plugins to extend the functionality of the KrakenD API Gateway. Happy coding!
