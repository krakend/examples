# KrakenD Flexible Configuration

Test environment to check [Flexible configuration](https://www.krakend.io/docs/configuration/flexible-config/) feature.

This repo contains a basic implementation for flexible configuration, including the use of variables, templates, code snippets and some basic logic to iterate over multiple endpoints.

### Contents:

- `krakend.json`: the base file, including calls to the variables, iteration over available endpoints and some code snippets.
- `partials/rate_limit_backend.tmlp`: a code snippet referenced from base file
- `settings/endpoint`: a collection of endpoints
- `settings/service.json`: basic configuration parameters for the service

## Running this test

### Using Docker

```shell
$ docker run \
--rm -it -p "8080:8080" \
-v "$PWD:/etc/krakend" \
-e FC_ENABLE=1 \
-e FC_SETTINGS=settings \
-e FC_PARTIALS=partial \
-e FC_OUT=out.json \
devopsfaith/krakend check -t -d -c "krakend.json"
```

### Using the binary locally

```shell
FC_ENABLE=1 \
FC_SETTINGS=settings \
FC_PARTIALS=partials \
FC_OUT=out.json \
krakend check -t -d -c "krakend.json"
```

Note: both alternatives will output a `out.json` file with the compiled version of the config file, useful for debugging purposes.
