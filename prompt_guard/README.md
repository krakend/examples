# KrakenD Prompt Guard Example

This is an example about how to use the [AI Gateway Prompt Guard](https://www.krakend.io/docs/enterprise/ai-gateway/prompt-guard/)
component.

Despite the fact this component is targeted at AI body payloads,
it can used on any `POST` / `PUT` endpoints.

## Run the example.

Start the environment with:

```
docker compose up -d
```

It will bring up the KrakenD instance runing the [`config/krakend/krakend.json`](./config/krakend/krakend.json)
config file, along with a fakepi server on port `8088`, and the other
services to view metrics and traces.

### The fake api

It exposes two endpoints:

- `http://localhost:8088/data/`: that serves the body used as a succesful
    response
- `http://localhost:8088/scorer/`: that randomly serves either a "block" or
    pass result.

There is a [test_fakeapi.sh](./client/test_fakeapi.sh) script under the client
directory to test those endpoints (that are used inside the krakend config).

### Run the tests

In the `client` directory, there is the `curl.sh` script that makes different
requests to see the behaviour of the prompt guards.

```bash
bash ./client/curl.sh
```

## The config

### Defined Guards.

At the service level `extra_config` we define a set of extra guards
that can be applied to endpoints:

- `email:` a regex to match emails (**not ready for production**, is only an example)
- `id_card`: a regex to match a simple type of ID numbers
- `random_scorer`: a guard that points to the fakeapi random scorer
- `failing_scorer`: a guard that points at not existing scorer, just to show
    the `on_failure_block` option that lets request pass if service is not
    available
- `max_size_policy`: an guard policy that checks the size of the body


### Endpoints

#### Default Prompt Guards

The `/defult_guards/` endpoint configuration example, with a list of empty guards,
that means only defaults will be applied.

In the `curl.sh` there are 3 requests made that shows how no poisoned data
can get the result back, while some other prompts / bodies will be
blocked by the default set of guards.

#### Regex Prompt Guard

There are 2 endpoints, for the custom regex guards, to showcase that
the filtering can be put at the endpoint level or at the backend level:

- `/personal/at_backend`
- `/personal/at_endpoint`

Both make use of the `email` and `id_card` regex guards defined
at the `"ai/prompt_guard"` service level

#### External classifier Prompt Guard

- `/classifier/random`

This endpoint has default prompt guards disabled, and uses two custom
classifiers:

- `failing_scorer`: a non existing url one with the flag of `on_failure_block` set
to true, to showcase that if it cannot connect it lets the request pass,
and
- `random_scorer`: that randomly blocks or lets the prompt pass

#### Policy Prompt Guard

This endpoint blocks in case of size of the payload is too big.

- `/policy/max_size`


## Observability



Traces can be checked at [http://localhost:16686](http://localhost:16686).

And metrics at [http://localhost:4000](http://localhost:4000) with user: `krakend`
and password `krakend`.

