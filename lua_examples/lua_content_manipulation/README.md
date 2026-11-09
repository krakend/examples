# Content manipulation: Lua vs JMESPath vs allow

This example returns the same result from three endpoints, each one using a
different technique to extract the `expandedBom` array from a large backend
response ([`data/bom.json`](./data/bom.json)) and return it as a collection:

| Endpoint          | Technique                                                                                               |
|-------------------|---------------------------------------------------------------------------------------------------------|
| `/using-lua`      | A [Lua script](https://www.krakend.io/docs/endpoints/lua/) ([`filter.lua`](./filter.lua)) deletes the unwanted keys, and `mapping` renames `expandedBom` to `collection` |
| `/using-jmespath` | A [JMESPath](https://www.krakend.io/docs/enterprise/endpoints/jmespath/) expression `expandedBom[*]`   |
| `/using-allow`    | The backend [`allow`](https://www.krakend.io/docs/backends/data-manipulation/) list keeps only `expandedBom`, and `mapping` renames it |

All three use `"output_encoding": "json-collection"` to return the array as
the root of the response.

## Run the example

You need a KrakenD Enterprise license, because JMESPath is an Enterprise
feature. Replace the placeholder `LICENSE` file in this folder with yours, and
start the stack:

```bash
docker compose up
```

Then compare the three responses:

```bash
curl http://localhost:8080/using-lua
curl http://localhost:8080/using-jmespath
curl http://localhost:8080/using-allow
```
