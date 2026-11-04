# Workflows Example

This example uses [Workflows](https://www.krakend.io/docs/enterprise/endpoints/workflows/)
to build a backend for frontend (BFF) for a fictional streaming platform.
A single endpoint returns everything the home screen needs, combining
parallel and sequential calls that a regular endpoint cannot mix.

## Requirements

- Docker and Docker Compose
- A KrakenD Enterprise license saved as `config/krakend/LICENSE` (need a
  trial license? [Contact us](https://www.krakend.io/contact-sales/))

## Run the example

Start the environment:

```bash
docker compose up -d
```

This starts KrakenD Enterprise on [http://localhost:8080](http://localhost:8080),
and a fake API on [http://localhost:8088](http://localhost:8088) that serves
the JSON files in [`config/fakehttpapi/data`](./config/fakehttpapi/data).

When you are done, stop everything with `docker compose down`.

## The `/home_dashboard` endpoint

The endpoint models the first screen a customer sees. It accepts the
`customer` query string, and optionally a `profile`. When there is no
`profile`, the dashboard uses the default profile of the customer.

Get the dashboard for the default profile:

```bash
curl "http://localhost:8080/home_dashboard?customer=1"
```

And for a specific profile:

```bash
curl "http://localhost:8080/home_dashboard?customer=1&profile=yy"
```

The response has three sections, which KrakenD fetches in parallel:

```json
{
  "header": { "movies": [ ... ] },
  "custom_content": {
    "profile": { "name": "Bart Sink Son", "profile_id": "xx" },
    "recommended": {
      "movies": [ ... ],
      "actors": [ ... ]
    }
  },
  "footer": { "movies": [ ... ] }
}
```

The `header` and `footer` are regular backends. The `custom_content` is a
workflow, because it needs the profile before it can fetch the recommended
content.

## How the workflow works

```
/home_dashboard
├── header                       (parallel)
├── custom_content workflow      (parallel, its backends run sequentially)
│   ├── 1. find_profile_id workflow
│   │   ├── /default_profile             only when no profile is passed
│   │   └── profile from the query string only when a profile is passed
│   └── 2. for_profile_content workflow  uses the profile_id from step 1
│       ├── /profile_recommended_movies/{profile_id}
│       └── /profile_recommended_actors
└── footer                       (parallel)
```

1. **Find the profile.** The `find_profile_id` workflow runs two backends,
   but only one of them returns data:
   - `/default_profile` has a [CEL](https://www.krakend.io/docs/endpoints/common-expression-language-cel/)
     check that skips it when the request includes a `profile`.
   - The other backend uses a
     [response body generator](https://www.krakend.io/docs/enterprise/backends/response-body-generator/)
     with the [`profile_id.tmpl`](./config/krakend/profile_id.tmpl) template,
     which builds the profile from the `profile` query string, or returns an
     empty object when there is none.
2. **Fetch the content of the profile.** The `custom_content` workflow is
   sequential and propagates the `profile` field of the first response as
   `resp0_profile`. The `for_profile_content` workflow then uses
   `{resp0_profile.profile_id}` to request the movies of that profile, and
   requests the recommended actors in parallel.

When you pass a `profile`, the KrakenD logs show
`request aborted by evaluator #0` for the `find_profile_id` workflow. That is
the CEL check skipping the default profile call, and it is expected: the
workflow sets `ignore_errors`, so it continues with the rest of the backends.
