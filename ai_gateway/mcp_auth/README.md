# KrakenD AI Gateway: MCP Auth

This demo showcases how MCP (Model Context Protocol) authorization works with KrakenD.

An MCP server allows AI agents to access external tools and data sources through a standardized protocol. 

The latest MCP spec states that the authorization server discovery can be
done either with a `WWW-Authenticate` header response, or a well known url
exposing the metadata. Clients **MUST** support both methods, so an MCP server
can use only one. 

This example shows how to display the authorization service metadata by 
using a `backend/static-filesystem` configuration: 

```json
{
  "endpoint": "/.well-known/*",
  "method": "GET",
  "timeout": "15s",
  "output_encoding": "no-op",
  "backend": [
    {
      "url_pattern": "/",
      "host": [ "http://ignore" ],
      "extra_config": {
         "backend/static-filesystem": {
            "path": "/etc/krakend/well_known"
         }
      }
    }
  ]
}
```

and under the `./config/krakend/well_known` directory, we have the 
`oauth-protected-resource` dir, with the `auth_mcp` file:

```
{
    "resource": "https://krakend_ee/auth_mcp",
    "authorization_servers": [
        "https://keycloak/realms/krakend/"
    ],
    "bearer_methods_supported": ["header"],
    "scopes_supported": [
        "basic", 
        "service_account",
        "roles"
    ]
}
```

that corresponds to the configured `/auth_mcp` endpoint in the configuration:

```json
{
      "endpoint": "/auth_mcp",
      "method": "POST",
      "timeout": "15s",
      "backend": [
        {
          "url_pattern": "/ignore",
          "host": [
            "http://ignore"
          ]
        }
      ],
      "extra_config": {
        "ai/mcp": {
          "server_name": "country-weather"
        },
        "auth/validator": {
          "alg": "RS256",
          "audience": [
          ],
          "disable_jwk_security": true,
          "jwk_url": "http://keycloak:8080/realms/krakend/protocol/openid-connect/certs",
          "roles": [
            "moderator"
          ],
          "roles_key": "realm_access.roles",
          "operation_debug": true,
          "roles_key_is_nested": true
        }
      }
    }
```

However, the well known url could be served from an nginx proxy directly
put in front of KrakenD if there is already that service in place.

Also, the use of authorization requires secure (HTTPS) urls, so in this example
we show how to generate self-signed certificates, and we put an nginx 
service in front of KrakenD to deal with TLS termination. 

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
cd ..
```

Then, add the created CA to the trusted certificates folder, and update the 
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

4. Start services:
```bash
docker-compose up -d
```

5. Test the protected MCP server by adding it to your MCP client.

For example, for Claude Code, you can run:

```
claude mcp add --transport http krakend_auth_example https://krakend_ee/auth_mcp 
```


### Warning

Claude uses offline tokens. In order to use them, the user must have the
`offline_access` role, and the client must have the `offline_access` scope.

The `/auth_mcp` endpoint only accepts users with the `moderator` realm role.
The imported realm includes the `sarahconnor` user, which has both the
`moderator` and `offline_access` roles.


## How It Works

**MCP Tool Definition**: The `get_country_info` tool is defined with its input schema and workflow,
and replies with a fake response from the `/fakeapi/*` defined endpoint

## Configuration

The MCP server configuration is in `config/krakend/krakend.json`.

The MCP endpoint definition includes:
- Server metadata (name, title, version, instructions)
- Tool definitions with input schemas
- Workflow configuration

## Resources

- [MCP Server Configuration](https://www.krakend.io/docs/enterprise/ai-gateway/mcp/)
