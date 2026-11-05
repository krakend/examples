# MCP Server with Dynamic Routing by Host

This example shows how a single [KrakenD MCP Server](https://www.krakend.io/docs/enterprise/ai-gateway/mcp-server/)
tool can reach a different backend depending on the hostname the client used to
connect, by combining the tool `input_headers` with
[Dynamic Routing](https://www.krakend.io/docs/enterprise/endpoints/dynamic-routing/).

The MCP server exposes one tool, `get_fantasy_planet`. Its response depends on
the `Host` header of the request:

- `Host: vegeta` returns the data of the Vegeta planet.
- `Host: namek` returns the data of the Namek planet.

This pattern is useful when several tenants or environments share the same
KrakenD MCP Server and each one must reach its own backend.

## Requirements

- Docker and Docker Compose
- A KrakenD Enterprise license saved as `LICENSE` in this directory (need a
  trial license? [Contact us](https://www.krakend.io/contact-sales/))

## Run the example

Start the environment:

```bash
docker compose up -d
```

This starts two services:

| Service    | URL                                            | Purpose                                                               |
|------------|------------------------------------------------|-----------------------------------------------------------------------|
| KrakenD EE | [http://localhost:8080](http://localhost:8080) | MCP server, with its entrypoint at `/fantasy`                         |
| nginx      | [http://localhost:8085](http://localhost:8085) | Backend with one virtual host per planet (`vegeta` and `namek`)       |

Then call the tool once per planet:

```bash
bash ./test_mcp.sh
```

Each call returns a different planet, even though the MCP request is the same:

```
--- Host: vegeta
event: message
data: {"jsonrpc":"2.0","id":1,"result":{"content":[{"type":"text","text":"{\"distance\":52939,\"planet\":\"Vegeta\"}"}],...}}

--- Host: namek
event: message
data: {"jsonrpc":"2.0","id":1,"result":{"content":[{"type":"text","text":"{\"distance\":293114,\"planet\":\"Namek\"}"}],...}}
```

When you are done, stop everything with `docker compose down`.

### Using a real MCP client

MCP clients usually do not let you set the `Host` header, but they send the
hostname of the URL you connect to. Add the planets to your `/etc/hosts`:

```
127.0.0.1 vegeta namek
```

Then connect your MCP client to `http://vegeta:8080/fantasy` or
`http://namek:8080/fantasy`. The `Host` header then includes the port
(`vegeta:8080`), and the example handles both forms.

## How it works

The tool declares which client headers it receives:

```json
"input_headers": ["Host"]
```

Its workflow then uses the header to build the backend URL:

```json
{
  "url_pattern": "/{input_headers.Host}/index.json",
  "host": ["http:/"],
  "disable_host_sanitize": true
}
```

As of KrakenD EE 3.0.3, Dynamic Routing variables are not replaced inside the
`host` field of an MCP workflow, so the example puts the whole destination in
`url_pattern`: the host
`http:/` plus the pattern `/vegeta/index.json` produce
`http://vegeta/index.json`. `disable_host_sanitize` keeps KrakenD from
rewriting that unusual host value.

On the Docker side, the nginx service has the network aliases `vegeta` and
`namek`, so both names resolve to it from the KrakenD container. nginx listens
on ports `80` and `8080`, so the backend URL works whether or not the `Host`
header includes the port, and it serves a different `index.json` for each
`server_name`.

## Security considerations

In this example the client decides which host KrakenD connects to. In a
production setup, never route on an unvalidated `Host` header: a client could
point KrakenD to any host reachable from the gateway. Make sure only the
expected values reach the tool, for instance by validating the header before
the request reaches KrakenD, or by restricting the names the gateway can
resolve.
