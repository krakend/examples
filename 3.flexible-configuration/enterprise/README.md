## Extended Flexible Configuration example
Example of extended flexible configuration (EFC) templating. This example is not here as a best-practices demonstration, it shows an implementation to demonstrate how to configure EFC using a **certain level of complexity**. Use this example for learning, but always choose the simplest way of managing the configuration for the organizational complexity you have. If you have a simple API, a single `krakend.json` is the best choice, and grow from there.

See the [EFC documentation](https://www.krakend.io/docs/enterprise/configuration/flexible-config/)

## Starting the demo
To test this configuration you need:

- Docker
- `cd` to this folder ($PWD)
- A `LICENSE` file

You can test this API locally with:

```bash
docker run --rm -it -p "8080:8080" \
    -e "KRAKEND_LICENSE_PATH=/LICENSE" \
    -e "FC_CONFIG=flexible_config.json" \
    -v "$PWD/LICENSE:/LICENSE:ro" \
    -v "$PWD:/etc/krakend" krakend/krakend-ee:watch
```

You can save files and the API is reloaded on the fly. The generated file is under `out.json`

This is an illustrative example mixing templates, and while maintining one file per endpoint is not necessarily the best choice, it demonstrates how to work with templates.

Play a little bit with it and let us know any questions

## One file = One endpoint. One repo/folder = One team
There are as many ways of organizing a KrakenD configuration as companies using KrakenD. While a simple approach usually works best, in this example, teams are supposed to manage their APIs editing the contents under `settings/endpoints/team-folder`. Each team has a different folder, which could be actually hosted in different repos and aggregated on build time. 

There are two teams in this example (named after famous 80s TV Shows):

- `A-team`
- `Knight-Rider`

The endpoints are declared as a single file, and for demonstration purposes they are declared in YAML, JSON or TOML formats. In this demo, teams manage endpoints by adding new files under their team folder. For instance, create now a new file, with any name, and paste the following:

```json
{
    "endpoint": "/foo"
}
```
This creates a configuration with the defaults (check this endpoint under `out.json`). The templates can be easily expanded, and to demonstrate this, the templates support a few more logic:

```json
{
    "endpoint": "/first",
    "url_pattern": "/api/v1/legacy/whatever",
    "method": "POST",
    "payload": {
        // JSON Schema validation
    },
    "jwt": ["admin","user","editor"],
    "custom_host": "http://custom-host.example.com"
}
```

## Changing environments
You will find a folder `environment` that contains settings that are specific to any environment (`common`) or overrides for specific environment.
In addition, the `flexible_config.json` loads the settings from the `development` environment. You can change this, but you can also pass a different file when starting KrakenD using the `FC_CONFIG` environment var.



