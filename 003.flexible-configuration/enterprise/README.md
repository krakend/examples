## Extended Flexible Configuration example
This is an example of extended flexible configuration (EFC) templating. It is not a best-practice demonstration; it shows an implementation of configuring EFC using a **certain level of complexity**. 

Use this example for learning, but always choose the simplest way of managing the configuration for your organizational complexity. If you have an uncomplicated API, a single `krakend.json` is the best choice, and then you grow from there.

See the [EFC documentation](https://www.krakend.io/docs/enterprise/configuration/flexible-config/)

## Starting the demo
To test this configuration, you need:

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

When you save files, the API is reloaded on the fly. The generated file is under `out.json`

This is an illustrative example of mixing templates. While **maintaining one file per endpoint is not necessarily the best choice**, it demonstrates how to work with templates.

Play a little bit with it, and get in touch if you have any questions

## One file = One endpoint. One repo/folder = One team
There are as many ways of organizing a KrakenD configuration as companies using KrakenD. While a simple approach usually works best, in this example, teams are supposed to manage their APIs **without knowing about KrakenD**.

The idea after this approach is that teams can manage APIs by declaring a metadata file with any format you find convenient. This metadata is used to populate the final API contract.

Teams should edit the contents under their `settings/endpoints/team-folder`. Each team folder could be hosted in different repos and aggregated during build time. 

You will find two teams in this example (named after famous 80s TV Shows):

- `A-team`
- `Knight-Rider`

The endpoints underneath are declared as a single file; for demonstration purposes, they are declared in YAML, JSON, or TOML formats. In this demo, teams manage endpoints by adding new files under their team folder. For instance, create now a new file with any name, and paste the following:

```json
{
    "endpoint": "/foo"
}
```
This creates a configuration with the defaults (check this endpoint under `out.json`). The templates can be easily expanded, and to demonstrate this, the templates support a few more logic, like adding payload validation, requiring JWT roles, or using non-default hosts:

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

## Code organization
The code tree has the following meaning:

```
├── environment
│   ├── common <-- contains all settings for all environments
│   │   └── infra.yaml <-- contains infrastructure values
│   ├── development <-- The development environment
│   │   └── infra.yaml <-- Overwrites some values of the common folder
│   └── production <-- The production environment
│       └── infra.yaml
├── flexible_config.json <-- Defines where to take the data from
├── krakend.json <-- The initial template
├── partials
├── settings
│   ├── endpoints <-- All the endpoints under this gateway
│   │   ├── A-team <-- This folder belongs to a specific team
│   │   │   ├── endpoint1.json <-- Below teams, there are the endpoints.
│   │   │   └── endpoint2.yml
│   │   └── Knight-Rider
│   │       ├── devon.toml
│   │       └── michael.json
│   └── service_extra_config.json
└── templates <-- The templates definition that render the final config
    ├── infra_endpoints.tmpl <-- endpoints specific to the Infra team
    ├── jwt_validation.tmpl <-- A template to validate JWT
    ├── payload_validation.tmpl <-- A template to validate JSON validates
    └── teams_endpoints.tmpl <-- A template to render all team's endpoints
```
## Changing environments
You will find a folder `environment` that contains settings specific to any environment (`common`) or overrides for a particular environment.

In addition, the `flexible_config.json` loads the settings from the `development` environment. You can edit this file and set another value, you can replace it with a placehodler, or you can also pass a completely different file when starting KrakenD using the `FC_CONFIG` environment var.