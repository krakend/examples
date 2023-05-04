# Merge and groupBy multiple sources with KrakenD

This configuration file defines an API Gateway using KrakenD, with two main endpoints exposed. Below is a brief description of each endpoint and the logic applied to them.

## Endpoints
1. `/components`: returns a JSON collection, merging and applying a group-by logic to the source data structure recovered from `/__internal-endpoint-raw-components` endpoint. The groupBy logic is implemented via a Lua script, `groupby.lua`.<br><br>

2. `/__internal-endpoint-raw-components`: consolidates data from two different sources (`/bom.json` and `/component.json`). The aggregation and flattening of data is done using `modifier/jmespath` plugin, that applies a JMESPath expression `[bom,component][]`.

### The content manipulation workflow

This workflow illustrate the content manipulation performed by this KrakenD example:

Retrieve data from two different collections:

```
Collection 1 (bom.json)        Collection 2 (component.json)
------------------------        -----------------------------
{                              {
"id": 1,                         "id": 1,
"name": "component1",            "metadata": "metadata1",
...                              ...
}                              }
{                              {
"id": 2,                         "id": 2,
"name": "component2",            "metadata": "metadata2",
...                              ...
}                              }
...
```

Combine and flatten the two collections into a single JSON array using the JMESPath expression [bom,component][]:

```
Aggregated and Flattened JSON Collection
----------------------------------------
[   
   {"id": 1, "name": "component1", ...},   
   {"id": 2, "name": "component2", ...},   
   {"id": 1, "metadata": "metadata1", ...},   
   {"id": 2, "metadata": "metadata2", ...},   
   ...
]
```
Apply the Lua script groupby.lua to aggregate fields from items with the same ID into a single JSON collection:

```
Final JSON Collection
---------------------
[
{"id": 1, "name": "component1", "metadata": "metadata1", ...},
{"id": 2, "name": "component2", "metadata": "metadata2", ...},
...
]
```

This workflow demonstrates how the API Gateway combines and manipulates data from two different collections into a single JSON collection, where each item includes all the fields from the original items with the same ID.
