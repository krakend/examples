# KrakenD API Gateway - Custom Rate Limiting PoC

## Overview
This repo tries to implement custom rate limits for different user tiers in the KrakenD API Gateway, using API keys. The KrakenD configuration file demonstrates the usage of tiered ratelimits introduced in the [2.7 enterprise version](https://krakend.io/blog/krakend-ee-2.7-release-notes).

### Configuration File Description
The KrakenD configuration file, `krakend.json`, is tailored for applying different rate limits based on user tiers (gold and silver) when using API keys.

#### Key Components:

- **auth/API-Keys**: Defines API keys for gold and silver user tiers (roles).
- `qos/ratelimit/tiered`: to define the ratelimit to apply to each defined tier.
- **endpoints**: Configures the main endpoint `/test` and the `qos/ratelimit/tiered`.

#### Notes:

Take into account that the role will be the first that matches the list of provided roles 
inside the endpoint's `extra_config` -> `auth/api-keys`. 

### Setup Instructions
1. **Install & Run KrakenD using Docker**:
   ```shell
   docker run -p "8080:8080" -v "$PWD:/etc/krakend" krakend/krakend-ee:2 run -dc krakend.json
   ```
   Ensure Docker is installed and the `krakend.json` file is in the current working directory.

### Testing the Configuration
To test the rate limiting, use the following curl commands:

- **For Gold Tier API Key**:
  ```shell
  curl -H'Authorization: Bearer gold-4d2c61e1-34c4-e96c-9456-15bd983c50' http://localhost:8080/test
  ```

- **For Silver Tier API Key**:
  ```shell
  curl -H'Authorization: Bearer silver-4d2c61e1-34c4-e96c-9456-15bd983c50' http://localhost:8080/test
  ```

Observe the rate limiting behavior based on the user tier.

#### Us the `hey` tool to use parallel requests

Download the [hey tool from github](https://github.com/rakyll/hey) and edit the `hey.sh`
bash script to update the env var with the path of the `hey` executable.

Then run the hey script to execute requests in parallel.
