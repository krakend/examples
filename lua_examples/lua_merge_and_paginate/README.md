# Aggregation, flattening and pagination example

The main goal of this configuration is to merge and flatten data from two different services and paginate the results.

## Endpoints

1. `/list/{page}`<br>
   This endpoint takes a page number as a parameter and serves paginated results of the merged data from the internal `/__internal-endpoint-list` endpoint. The output encoding is set to no-op, which means that the response is not encoded or transformed.<br><br>The modifier/lua-proxy extra config is used to run a Lua script (`paginate.lua`) that takes care of merging the data and paginating it. The live option is set to true, which allows the Lua script to be updated at runtime without restarting the service.<br><br>The response body of the endpoint is the result of the `paginate(req, resp)` function from the Lua script.<br><br>

2. `/__internal-endpoint-list`<br>
   This internal endpoint retrieves data from two services, `/agents.json` and `/users.json`, and merges the results using the modifier/jmespath extra config. The expr field contains the JMESPath expression used to merge the data into a single users_list array.<br><br>The `plugin/req-resp-modifier` extra config is used to apply an IP filter, named `ip-filter`. The filter runs in allow mode with a list of private ranges (`192.168.0.0/24`, `172.0.0.0/8`, `10.0.0.0/16`) plus `127.0.0.1` and `::1`, so only requests from the local machine or the Docker network can access this internal endpoint.

The service also loads the `url-rewrite` plugin (`plugin/http-server`), which rewrites `/list` and `/list/` to `/list/1`.


You can run this example by executing a `docker compose up`
