# KrakenD AI Gateway: MCP Server Demo

This demo showcases KrakenD as an MCP (Model Context Protocol) server that provides unified access to multiple external APIs. It demonstrates how KrakenD can act as a tool provider for AI agents by aggregating REST and GraphQL endpoints into a single, AI-friendly interface.

An MCP server allows AI agents to access external tools and data sources through a standardized protocol. This example shows how KrakenD can orchestrate multiple API calls, transform data, and expose it as MCP tools.

## What This Demonstrates

**MCP Server Implementation**: KrakenD exposes an MCP endpoint that AI agents can use to retrieve comprehensive country information and weather data.

**Key capabilities**:
- MCP server configuration with tool definitions
- Sequential proxying to aggregate multiple APIs (REST + GraphQL)
- Data propagation between backend calls
- Response transformation using JMESPath and Lua
- Unified response from heterogeneous data sources

## Quick Start

### Prerequisites

- Docker and Docker Compose
- KrakenD Enterprise license (need a trial license? [Contact us](https://www.krakend.io/contact-sales/))

### Setup

1. Add your KrakenD Enterprise license as `LICENSE` in the root directory

2. Start services:
```bash
docker-compose up -d
```

3. Test the MCP server by calling the endpoint:
```bash
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "method": "tools/call",
    "params": {
      "name": "get_country_info",
      "arguments": {
        "country_code": "ES"
      }
    },
    "id": 1
  }'
```

## How It Works

1. **MCP Tool Definition**: The `get_country_info` tool is defined with its input schema and workflow
2. **Sequential Backend Calls**:
   - First: REST Countries API fetches geography, population, borders, and flag data
   - Second: GraphQL Countries API retrieves currency, languages, and emoji
   - Third: Open-Meteo Weather API gets current weather for the capital city
3. **Data Propagation**: Capital coordinates from the first call are passed to the weather API
4. **Response Aggregation**: All data is merged into a unified response using JMESPath
5. **Lua Processing**: Custom Lua script flattens capital coordinates for easier access

## Configuration

The MCP server configuration is in `config/krakend/krakend.json`. Lua transformations are in `config/krakend/lua/`.

The MCP endpoint definition includes:
- Server metadata (name, title, version, instructions)
- Tool definitions with input schemas
- Workflow configuration with backend orchestration

## Additional Features

**Sequential Proxying**: Demonstrates how to chain API calls and propagate data between them

**Multi-Protocol Aggregation**: Combines REST and GraphQL APIs in a single workflow

**Error Handling**: Returns error messages when tool execution fails

## Resources

- [KrakenD AI Gateway Documentation](https://www.krakend.io/docs/enterprise/ai-gateway/)
- [MCP Server Configuration](https://www.krakend.io/docs/enterprise/ai-gateway/mcp/)
- [Sequential Proxying](https://www.krakend.io/docs/endpoints/sequential-proxy/)
- [GraphQL Integration](https://www.krakend.io/docs/backends/graphql/)
