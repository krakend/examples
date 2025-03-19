# gRPC

**Enterprise only feature** 

Place your license file under the `config/krakend` directory before starting 
docker compose.

You must have installed the [Protocol Buffer Compiler](https://grpc.io/docs/protoc-installation/) 
to have `protoc` command available.

Then run:

```
$ make docker_up
```

Use the clients under the `client` dir to test different request to the
krakend service.


## Endpoints

### `/grpc/travel`

#### features

- **use query strings as input fields**: at the enpoint level we define that we
want to allow to let pass `lat` and `lon`, and those are later used in
`backend` -> `extra_config` -> `backend/grpc` -> `input_mapping` to fill
fields inside the request payload.

- **well known types**: for `time` and `duration` types, it accepts string
formatted inputs instead of using directly the gRPC payload fields of those 
types.

- **backend gRPC and HTTP**: each backend can be of its kind, and the merging
    grouping, and those kind features work as expected


The endpoint **require that at a service level** 
(under the `extra_config` -> `grpc` -> `catalog` section), you
provide a catalog of definitions. There you can list directories that
will be traversed looking for `.pb` files, or also add filenames. In case of 
conflict (finding different files with the same definition, the first found one
prevails).

In the backends we use `/flight_finder.Flights/FindFlights` and 
`/train_finder.Trains/FindTrains`. 

## gRPC services

### `flight_finder.Flights`

In order to expose services (using the same http port that for other endpoints),
we need to add them under `extra_config` -> `services`, in a list. There for 
each "service" you can expose a list of "methods" (that should already exist
in the catalog files). We can map some of the inner fields as "params" to be 
used later on by the backend. Backends can be the same than in the endpoints
section (there is nothing special for being called from a gRPC exposed method).

## The fake services

Under the `images` directory, there is the source code for two fake services,
that pretend to be a way to get information about trains and flights.

You can take a look at the [README](./images/grpc/README.md) there to get more info.

## The clients

- `clients`: here there are the scripts to make calls to krakend service
    using different tools: `curl`, `grpcurl`, ..
- `images`: here there is the source code for the fake APIs in gRPC
- `config`: the config files for the **KrakenD** service, and the fake
    api server
