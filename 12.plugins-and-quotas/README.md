# KrakenD Plugins and Quotas Example

This example demonstrates how to implement **token-based quota management** for streaming AI/LLM endpoints using KrakenD Enterprise Edition with a custom middleware plugin.

## Overview

The example showcases:

- **Custom Middleware Plugin**: A Go plugin that intercepts streaming responses and tracks token usage in real-time
- **Quota Management**: Multi-tier rate limiting (gold, silver, bronze) using KrakenD's governance/processors
- **Stream Processing**: Parsing streaming responses to extract token usage from the response body
- **Persistent Tracking**: Quota consumption tracked across requests with automatic storage management
- **Mock Backend**: An OpenAI-compatible mock server that streams responses with token metadata

## Architecture

```
┌──────────┐     ┌─────────────────────────────┐     ┌──────────┐
│  Client  │────▶│  KrakenD EE                 │────▶│ Backend  │
└──────────┘     │  - Quota Pre-check          │     │ (Mock    │
                 │  - Stream Wrapper           │     │  OpenAI) │
                 │  - Token Extraction         │     └──────────┘
                 │  - Quota Update             │
                 │                             │
                 │  Quota Processor handles    │
                 │  storage transparently      │
                 └─────────────────────────────┘
```

## Components

### 1. KrakenD Configuration (`config/krakend/krakend.json`)
- Defines quota rules for three tiers (gold, silver, bronze)
- Configures the quota processor with storage backend
- Sets up endpoint with the custom middleware plugin

### 2. Middleware Plugin (`plugins/quota-control-mw/middleware.go`)
- **Pre-check**: Validates if user has quota before forwarding request
- **Stream Wrapping**: Intercepts the response stream from the backend
- **Token Extraction**: Uses regex to parse token usage from streaming chunks
- **Quota Update**: Dynamically updates quota based on actual token consumption

### 3. Mock Backend (`backend/openai-mock/main.go`)
- Simulates OpenAI's streaming API format
- Returns Server-Sent Events (SSE) with mock responses
- Includes token usage metadata in the final chunk

## Prerequisites

- Docker and Docker Compose
- KrakenD Enterprise Edition license (place `LICENSE` file in the root directory)

## Running the Example

1. **Ensure you have a KrakenD EE license file** named `LICENSE` in this directory

2. **Start all services**:
   ```bash
   docker compose up --build
   ```

   This will start:
   - KrakenD EE on port `8080`
   - Mock backend on port `8090`
   - Supporting services (quota storage)

3. **Test the endpoint**:
   ```bash
   curl -X POST http://localhost:8080/ \
     -H "Content-Type: application/json" \
     -d '{
       "messages": [{"role": "user", "content": "Hello!"}]
     }'
   ```

   You should see a streaming response with the mock data.

## How It Works

1. **Request arrives** at KrakenD's `/` endpoint
2. **Middleware pre-checks** quota limits (weightless check with 0 tokens)
3. If quota allows, **request is forwarded** to the backend
4. **Backend streams** the response in OpenAI-compatible SSE format
5. **Middleware wraps** the response stream with a custom reader
6. As **chunks are read**, the plugin:
   - Searches for usage metadata in the stream
   - Extracts token counts using regex pattern
   - Updates the quota processor with actual consumption
7. **Client receives** the full streaming response
8. **Quota is accurately tracked** based on real token usage

## Quota Tiers

The configuration defines three quota tiers:

| Tier   | Hourly Limit | Daily Limit |
|--------|--------------|-------------|
| Gold   | 1,000 tokens | 5,000 tokens|
| Silver | 500 tokens   | 2,000 tokens|
| Bronze | 200 tokens   | 1,000 tokens|

**Note**: In the plugin code (`middleware.go:73`), the tier is hardcoded to `"admin"` (which maps to the `gold` tier). In a production setup, you would extract this from:
- JWT claims
- Request headers (e.g., `X-User-Tier`)
- API key lookup
- Query parameters

## Key Features

### 1. **Streaming-Aware Quota Management**
Unlike traditional rate limiting that counts requests, this example counts tokens consumed during streaming, making it ideal for AI/LLM APIs where cost is token-based.

### 2. **Pre-check + Post-update Pattern**
- Pre-check (weight=0): Fast rejection of users who've already exceeded quota
- Post-update (weight=actual): Accurate quota deduction after response completes

### 3. **Regex-based Token Extraction**
The plugin uses a regex pattern to extract token usage from the streaming response:
```go
usagePattern: `{"prompt_tokens":(\d+),"completion_tokens":(\d+),"total_tokens":(\d+)}`
```

### 4. **Bloom Filter Optimization**
The configuration includes a rejecter cache (Bloom filter) to quickly deny previously blocked users without storage lookups.

## Customization

### Change User/Tier Extraction
Modify `middleware.go:73-74` to extract tier and user ID from the request:

```go
// Example: Extract from JWT claims
tier := extractFromJWT(reqw, "tier")
userId := extractFromJWT(reqw, "user_id")

// Example: Extract from headers
tier := reqw.Headers()["X-User-Tier"][0]
userId := reqw.Headers()["X-User-Id"][0]
```

### Adjust Quota Limits
Edit `config/krakend/krakend.json` in the `governance/processors.quotas.rules` section.

### Change Token Pattern
If your backend uses a different format for token usage, update the regex pattern in `middleware.go:62`.

## Building the Plugin Manually

If you want to build the plugin outside Docker:

```bash
cd plugins/quota-control-mw
make go.mod  # Initialize Go module
make amd64   # Build for AMD64
# or
make arm64   # Build for ARM64
```

## Troubleshooting

**Plugin not loading**: Check KrakenD logs for plugin-related errors. Ensure the `.so` file is in `/opt/krakend/plugins/` inside the container.

**Quota not working**: Verify all services are running with `docker compose ps`. Check KrakenD logs for quota processor initialization errors.

**Tokens not being tracked**: Check that the token usage format in the backend response matches the regex pattern in the plugin.

## License

This example requires a valid KrakenD Enterprise Edition license. Place your `LICENSE` file in the root directory.

## References

- [KrakenD Plugins Documentation](https://www.krakend.io/docs/enterprise/extending/)
- [KrakenD Quota Management](https://www.krakend.io/docs/governance/quota/)
- [OpenAI Streaming Format](https://platform.openai.com/docs/api-reference/streaming)
