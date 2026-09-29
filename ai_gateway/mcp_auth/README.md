# KrakenD AI Gateway: MCP Auth

This demo showcases KrakenD as an MCP (Model Context Protocol) authorization works.

An MCP server allows AI agents to access external tools and data sources through a standardized protocol. This example shows how KrakenD can orchestrate multiple API calls, transform data, and expose it as MCP tools.

## Quick Start

### Prerequisites

- Docker and Docker Compose
- KrakenD Enterprise license (need a trial license? [Contact us](https://www.krakend.io/contact-sales/))

### Setup

1. Add your KrakenD Enterprise license as `LICENSE` in the root directory

2. Create self signed certificates

```
cd certs
make all
cd ...
```

Then, add the created CA to the trusted certifificates folder, and update the 
certificates:

```
sudo cp ./certs/ca.crt /usr/local/share/ca-certificates/krakend_examples.crt
sudo update-ca-certificates
```

3. Edit your `/etc/hosts` 
   
To make requests using the host that matches the certificates for the 
`krakend_ee` instance and `keycloak` instance, we need the computer to
resolve those hostnames to `127.0.0.1` that is where our docker containers 
have the ports exposed:

```
127.0.0.1 krakend_ee keycloak
```

3. Start services:
```bash
docker-compose up -d
```

4. Test the protected MCP server, by adding it to your mcp client. 
   
For example, for claude, you can run:

```
claude mcp add --transport http krakend_auth_example https://krakend_ee/auth_mcp 
```


### Warning

Claude uses offline tokens, in order to be used, the user must have the 
`offline_token` role. Also the client must have the `offline_token` scope.


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
