# GRPC Example

In this directory you can find the source code used to build
gRPC backend services that can be used to test KrakenD features. 

The services are designed to model some trip finding services,
so we can put an example of an endpoint to search travel options
for a trip.

So we have:

- the `flights` service (built upon a `proto3` gRPC service definition).
- the `trains` service (built upon a `proto2` gRPC service definition).

## Building

To build the images run `make all_images`

### Folder structure:

- `contracts`: here is were we place all the `.proto` files definitions,
    for our two services, as well as the `.proto` files for the 
    dependencies our definitions use.
    
- `genlibs`: here is where the generated **client** and **server** 
    `protoc` generated code, to implement the services
    
- `flights`: the actual implementation of the flights service
    (that makes use of the generated lib under `genlibs/flights`).

- `trains`: the actual implementation of the trains service
    (that makes use of the genrated lib under `genlibs/trains`).

- `defs`: here we have the binary definitions of the `.proto`
    files.

## Shared library

Under the `contracts/lib` folder there are some `.proto`
definitions that are used in both services, to try
to mimic a real use case where an organization might
have some shared `gRPC` data definitions. 

Those definitions are:
    - `Page`: for requesting and sending paginated results
    - `GeoPosition`: just for grouping the `latitude` and 
        `longitude` in a single type.
    - `Address`: a human readable address
    - `Location`: a place that have an address and is
        geopositioned.
    - `TimeRange`: just a way to group two timestamps
        to define a time slice.


## Flights

The `flights` service is based on [a `proto3` gRPC service
definition](./contracts/flights/flights.proto)

The service exposes two `gRPC` calls:

- `FindFlight`: that is used to search for flights
- `BookFlight`: that is used to book for a flight

### `FindFlight` request

It uses several type from our own defined library to 
start a search: Page, Location and TimeRange to define
the page of results, the origin and time reange of
departure, and the destination and the time range of arrival.

There is also a `classes` array, so we can filter by the 
classes that we want to flight on (like `ECONOMY` or `BUSINESS`).

And also a filter requiring to just find flight that have 
a minimum discount, either by amount **OR** percentage (this
part does not make much sense, but is just to showcase the
usage of the `oneof`'s `proto3` feature.

### `BookFlight` request

This requests have some fields that cannot be mapped from
param / query strings, as it hits some limitations
explained in [gRPC passing user parameters](https://www.krakend.io/docs/enterprise/backends/grpc/#passing-user-parameters) :

- list of objects

#### Discounts

In order to test the usage of a `proto3`'s map as input,
we use the `Passenger`'s `discounts` field, that will
map a discount name, to its discount value to be applied.


## Trains

The `trains` service is based on a `proto2` gRPC service
definition.

The server uses a **self signed** certificate, to
showcase how a `gRPC` service can use `TLS`.

The main `Train` gRPC call has a couple of fields
declared as `required` (a `proto2` "feature" that
is no longer supported in `proto3`)



## Makefile

The `Makefile` contains the different steps used to build
the services. 

When a service is created or modified usually the steps to
follow are:

1. Modify the `contracts`, to update the interface the services
    will use.
2. Recreate the `grpc` librarie to be used by the services implementation.
3. Modify the services, and create the new binaries.
4. Dockerize the new service
5. Create the binary proto definitions to be used in KrakenD


Following these steps we document the `Makefile` targets to use:

### `get_known_types`

Before start writting our own contracts, we want to fetch
the "well known types" proto definitions, like the `Timestamp`
or the `Duration` types. 

Those types are "included" in the `protoc` tools, so when
generating your own contracts, you do not need to 
reference them explicitly, but if you want to generate
the binary definitions, you will need to have the `.proto` sources.

Run:

```
make get_known_types
```

to fetch them.

Once you have those types, you can write and modify your own ones.

### `generate_grpc_flights_lib` and `generate_grpc_trains_lib`

Those two targets will use the `protoc` tool to generate the `Go`
code for the `client` and `server` parts of the fake services,
ane will place them under the `genlib` directory.


## The flighs service: `flights_server`

This target will just build the gRPC service for flights.

```
make flights_server
```

## The trains service: `trains_server` and `generate_trains_cert`

The trains services is coded to make use of a certificate, so
there is an extra targeto to regenerate a self signed certificate
that will be placed under `trains/certs`.

```
make trains_server
make generate_trains_cert
```

### `bin_proto_multiple_files` and `bin_proto_single_file`

These are the steps to generate the binary `.pb` definitions
from the `.proto` source files. Those binary definitions
can be used for the KrakenD gRPC catalog.

You should **use only one of these two targets**: as it
is explained in the [Generating Binary Protocol Buffer Files](https://www.krakend.io/docs/enterprise/backends/grpc/#generating-binary-protocol-buffer-files-pb)
official KrakenD documentation, binary files can be each one
in its own file (`bin_proto_multiple_files`) or all definitions
can be collected in a single file (`bin_proto_single_file`).

The multiplle files ones makes use of a script found in the
`contracts` directory called `compile.sh`.

