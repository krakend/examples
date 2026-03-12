# Injecting Redis into KrakenD plugins

This example demonstrates how to build and use custom KrakenD Enterprise plugins with Redis integration. It includes a simple middleware plugin that increments a counter in Redis for each API request.

## What's Inside

- **KrakenD EE**: API Gateway with plugin support
- **Redis**: In-memory data store for the counter
- **Custom Plugin**: `counter-example-mw` - A Go plugin that increments a Redis counter on each request

## Prerequisites

- Docker
- Docker Compose

## Project Structure

```
.
├── docker-compose.yml              # Service orchestration
├── Dockerfile                      # Multi-stage build for KrakenD with plugin
├── config/
│   └── krakend/
│       └── krakend.json           # KrakenD configuration
└── plugins/
    └── counter-example-mw/        # Custom middleware plugin
        ├── middleware.go          # Plugin implementation
        ├── Makefile              # Build configuration
        ├── go.mod
        └── go.sum
```

## Quick Start

### 1. Launch the Stack

```bash
docker compose up --build
```

This will:
- Build the custom plugin as a shared object (.so file)
- Create a KrakenD EE container with the plugin loaded
- Start a Redis instance
- Expose KrakenD on port 8080 and Redis on port 6379

### 2. Test the API

Once the stack is running, test the endpoint:

```bash
curl http://localhost:8080/
```

Each request will:
- Pass through the custom middleware plugin
- Increment a counter in Redis with key `foo-counter:some-id`
- Return the response from the echo endpoint

### 3. Check the Redis Counter

Connect to Redis and check the counter value:

```bash
# Connect to Redis CLI
docker compose exec redis redis-cli

# Select database 1 (as configured in krakend.json)
SELECT 1

# Get the counter value
GET foo-counter:some-id

# Exit Redis CLI
EXIT
```

The counter should increment with each request to the API.

## Configuration Details

### KrakenD Configuration

The `config/krakend/krakend.json` file configures:

- **Redis Connection**: Points to the Redis service on port 6379, database 1
- **Plugin Loading**: Loads `.so` files from `/opt/krakend/plugins/`
- **Middleware**: Applies `counter-example-mw` plugin to the `/` endpoint
- **Debug Features**: Enables `echo_endpoint` and `debug_endpoint` for testing

### Plugin Configuration

The plugin accepts configuration in the endpoint's `extra_config`:

```json
"plugin/middleware": {
  "name": ["counter-example-mw"],
  "counter-example-mw": {
    "key_prefix": "foo-counter"
  }
}
```

The `key_prefix` is used to create Redis keys in the format: `{key_prefix}:some-id`

## How the Plugin Works

The `counter-example-mw` plugin:

1. Registers itself as a KrakenD middleware
2. Gets access to the Redis client configured as "local_instance"
3. On each request:
   - Passes the request to the next middleware/backend
   - Increments a counter in Redis using the configured key prefix
   - Returns the response

## Modifying the Plugin

### Change the Plugin Logic

1. Edit `plugins/counter-example-mw/middleware.go`
2. Rebuild the stack:

```bash
docker compose up --build
```

## Stopping the Stack

```bash
docker compose down
```

To also remove volumes (including Redis data):

```bash
docker compose down -v
```

## Troubleshooting

### Plugin Not Loading

Check the KrakenD logs for plugin loading errors:

```bash
docker compose logs krakend_ee
```

Look for messages like `[PLUGIN: counter-example-mw]` to verify the plugin loaded successfully.

### Redis Connection Issues

Verify Redis is running and accessible:

```bash
docker compose ps redis
docker compose exec redis redis-cli PING
```

Should return `PONG` if Redis is working correctly.

### Port Conflicts

If ports 8080 or 6379 are already in use, modify the port mappings in `docker-compose.yml`:

```yaml
ports:
  - "8081:8080"  # Change external port to 8081
```

## Additional Resources

- [KrakenD Plugin Documentation](https://www.krakend.io/docs/extending/)
- [Injecting Redis in plugins](https://www.krakend.io/docs/enterprise/extending/injecting-redis-in-plugins/)
- [KrakenD Redis Service configuration](https://www.krakend.io/docs/enterprise/service-settings/redis-connection-pools/)
- [Go Redis Client](https://github.com/redis/go-redis)
