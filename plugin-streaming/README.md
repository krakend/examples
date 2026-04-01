# KrakenD Plugins and Quotas Example

This example demonstrates how to implement streaming manipulation using KrakenD Enterprise Edition with a middleware plugin and a response modifier plugin.

## Overview

The example showcases how to manipulate stream responses. It does so by adding a `processed_time` key to the `data: {}` lines sent by the backend, which emulates a streaming session with OpenAI SSE response

## Running the Example

1. **Ensure you have a KrakenD EE license file** named `LICENSE` in the `config/krakend` directory

2. **build the plugin**

   ```bash
   make build
   ```

2. **Start all services**:

   ```bash
   docker compose up
   ```

   This will start:
   - KrakenD EE on port `8080`
   - Mock backend on port `8090`

3. **Test the endpoints**:

Middleware plugin:
   ```bash
   curl -X POST http://localhost:8080/middleware \
     -H "Content-Type: application/json" \
     -d '{
       "messages": [{"role": "user", "content": "Hello!"}]
     }'
   ```

Response modifier plugin:
   ```bash
   curl -X POST http://localhost:8080/response \
     -H "Content-Type: application/json" \
     -d '{
       "messages": [{"role": "user", "content": "Hello!"}]
     }'
   ```

## Troubleshooting

**Plugin not loading**: Check KrakenD logs for plugin-related errors. Ensure the `.so` is in `plugins/streaming-modifier` folder after running `make build`

**unsupported relocation type 7**: This is most likely a mismatch between your architecture and the one the plugin is built for. If running on arm64, change the Makefile to use `arm64` target instead of `amd64`

## License

This example requires a valid KrakenD Enterprise Edition license. Place your `LICENSE` file in the `./config/krakend` directory.

## References

- [KrakenD Plugins Documentation](https://www.krakend.io/docs/enterprise/extending/)
- [OpenAI Streaming Format](https://platform.openai.com/docs/api-reference/streaming)
