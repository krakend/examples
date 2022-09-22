# KrakenD Flexible Configuration

<div align="center">

| <a href="https://youtu.be/U1LHoKWy0HU"><img src="https://i.ytimg.com/vi/U1LHoKWy0HU/maxresdefault.jpg" width="300" heigth="300"></a><br>[Video in English](https://youtu.be/U1LHoKWy0HU) | <a href="https://youtu.be/qeTeLPLnkIY"><img src="https://i.ytimg.com/vi/qeTeLPLnkIY/maxresdefault.jpg" width="300" heigth="300"></a><br>[Video in Spanish](https://youtu.be/qeTeLPLnkIY) |
|---|---|

</div>

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
-e FC_PARTIALS=partials \
-e FC_OUT=out.json \
-e SERVICE_NAME="KrakenD API Gateway" \
devopsfaith/krakend check -t -d -c "krakend.json"
```

### Using the binary locally

```shell
FC_ENABLE=1 \
FC_SETTINGS=settings \
FC_PARTIALS=partials \
FC_OUT=out.json \
SERVICE_NAME="KrakenD API Gateway" \
krakend check -tdc "krakend.json"
```

Note: both alternatives will output a `out.json` file with the compiled version of the config file, useful for debugging purposes.
