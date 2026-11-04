# KrakenD Prompt Guard Example

This example shows how to use the [AI Gateway Prompt Guard](https://www.krakend.io/docs/enterprise/ai-gateway/prompt-guard/)
to inspect request bodies and block prompt injections, jailbreaks, personal
data, or any payload that breaks your own rules.

Although Prompt Guard is designed for AI payloads, it works on any endpoint
that receives a body (`POST`, `PUT`, ...), so you can use it in front of
regular APIs too.

## Requirements

- Docker and Docker Compose
- A KrakenD Enterprise license. Save it as `config/krakend/LICENSE` (need a
  trial license? [Contact us](https://www.krakend.io/contact-sales/))

## Run the example

Start the environment:

```bash
docker compose up -d
```

This starts the following services:

| Service      | URL                                              | Purpose                                                         |
|--------------|--------------------------------------------------|-----------------------------------------------------------------|
| KrakenD EE   | [http://localhost:8080](http://localhost:8080)   | The gateway, running [`config/krakend/krakend.json`](./config/krakend/krakend.json) |
| Fake API     | [http://localhost:8088](http://localhost:8088)   | Backend and fake classifier used by KrakenD                    |
| Jaeger       | [http://localhost:16686](http://localhost:16686) | Traces                                                          |
| Grafana      | [http://localhost:4000](http://localhost:4000)   | Dashboards (user `krakend`, password `krakend`)                 |
| Prometheus   | [http://localhost:9090](http://localhost:9090)   | Metrics storage                                                 |

When you are done, stop everything with `docker compose down`.

### The fake API

The fake API exposes two `POST` endpoints:

- `/data/`: returns a fixed JSON body, used as the backend response when a
  request is not blocked.
- `/scorer/`: a fake classifier that randomly returns a low score (`0.1`, the
  request passes) about 60% of the time, or a high score (`0.99`, the request
  is blocked) the other 40%.

You can call them directly with [`client/test_fakeapi.sh`](./client/test_fakeapi.sh).

### Send the test requests

The [`client/curl.sh`](./client/curl.sh) script sends a set of requests to
every endpoint, each one labeled with the expected outcome:

```bash
bash ./client/curl.sh
```

A blocked request returns an empty `500 Internal Server Error`, and the
KrakenD log shows `prompt guard blocked` for that endpoint.

## The configuration

Prompt Guard is configured in two places:

1. The service level `extra_config` declares your custom guards.
2. Each endpoint or backend lists, in its own `extra_config`, which guards
   apply to it.

KrakenD also ships a set of default guards. They are active on every endpoint
or backend that includes the `ai/prompt-guard` namespace, unless you disable
them with `!all_defaults` (or by severity or group, see the
[documentation](https://www.krakend.io/docs/enterprise/ai-gateway/prompt-guard/)).

### Custom guards

The service level declares these guards:

- `email` (`regex`): matches email addresses. The pattern is a simple example,
  not production ready.
- `id_card` (`regex`): matches a simple ID card number format.
- `random_scorer` (`classifier`): sends the body to the fake `/scorer/`
  endpoint and blocks the request when the score is `0.6` or higher.
- `failing_scorer` (`classifier`): points to a classifier that does not exist.
  Because `on_failure_block` is `false`, requests pass when the classifier
  cannot be reached.
- `max_size_policy` (`policy`): a CEL expression that blocks bodies bigger
  than 256 bytes.

Policy expressions describe what to **block**: a request is rejected when any
of the expressions evaluates to `true`.

### Endpoints

| Endpoint                  | Guards                                | What it shows                                                                    |
|---------------------------|---------------------------------------|----------------------------------------------------------------------------------|
| `/default_guards/`        | Defaults only                         | An empty `guards` list keeps only the built-in guards                            |
| `/personal/at_endpoint`   | `email`, `id_card`                    | Guards applied at the endpoint level                                             |
| `/personal/at_backend`    | `email`, `id_card`                    | The same guards applied at the backend level                                     |
| `/classifier/random`      | `random_scorer`, `failing_scorer`     | External classifiers, and what happens when one of them is unavailable           |
| `/policy/max_size`        | `max_size_policy`                     | A CEL policy that rejects large payloads                                         |

All endpoints except `/default_guards/` use `!all_defaults`, so only the
listed guards run.

Expected results when running `curl.sh`:

- `/default_guards/`: a harmless body passes, while a prompt asking to forget
  previous instructions and a jailbreak attempt are blocked.
- `/personal/*`: a body without personal data passes, while bodies containing
  an email address or an ID card number are blocked.
- `/classifier/random`: each request passes or is blocked at random,
  depending on the score returned by the fake classifier. The unreachable
  `failing_scorer` never blocks.
- `/policy/max_size`: a short body passes, while the long body in
  [`client/payload/samuel.json`](./client/payload/samuel.json) is blocked.

## Observability

Every time a prompt guard runs, KrakenD records it in both metrics and traces:

- **Metrics**: the `krakend_promptguard_check_requests_total` counter in
  Prometheus, with the `http_route` and `krakend_promptguard_blocked`
  (`true` or `false`) labels. Grafana includes a **KrakenD - Prompt Guard**
  dashboard with these numbers, next to the general KrakenD dashboard.
- **Traces**: open Jaeger, select the `krakend_prompt_guard` service, and look
  for the `krakend.promptguard.blocker` attribute, which tells you which guard
  (or default guard group) blocked the request.
