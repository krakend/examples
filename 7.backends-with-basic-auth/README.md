# KrakenD Basic Authentication Example

This example shows how to connect KrakenD with backends that require Basic
Authentication, using a [Martian modifier](https://www.krakend.io/docs/backends/martian/)
in the backend configuration to inject the `Authorization` header.

The `fake_api` service is a small static server that protects
`/protected-response.json` with the user `krakend` and the password
`password` (see [`data/httpd.conf`](./data/httpd.conf)).

## Run the example

```bash
docker compose up -d
```

Then call the two endpoints:

```bash
curl -i http://localhost:8080/test-basic-auth-ok
curl -i http://localhost:8080/test-basic-auth-ko
```

- `/test-basic-auth-ok` adds the `Authorization` header with the right
  credentials, so the backend returns the protected JSON.
- `/test-basic-auth-ko` calls the same backend without credentials. The
  backend answers `401`, and KrakenD returns an error.

Stop everything with `docker compose down`.
