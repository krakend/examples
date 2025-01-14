# KrakenD API Gateway - Custom Rate Limiting PoC

## Overview
This repo tries to implement custom rate limits for different user tiers in the KrakenD API Gateway, using API keys. The KrakenD configuration file demonstrates a possible method, using Lua and internal endpoints, to achieve this functionality.

### Configuration File Description
The KrakenD configuration file, `krakend.json`, is tailored for applying different rate limits based on user tiers (gold and silver) when using API keys.

#### Key Components:
- **auth/API-Keys**: Defines API keys for gold and silver user tiers.
- **endpoints**: Configures the main endpoint `/test` and two internal endpoints for gold and silver tiers.

#### Method for Implementing Custom Rate Limits:
**URL Modification with Lua Scripting**: Involves modifying the URL by appending the user tier at the end using Lua scripting at the backend level.

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
