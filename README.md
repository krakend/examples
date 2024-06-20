# KrakenD Examples Repository

Welcome to the KrakenD Examples repository. This repository contains various examples demonstrating different features and configurations of the KrakenD API Gateway. Below you will find a list of the examples contained in this repository, along with a brief description and links to each example.

## Examples

### [0.krakend-docker](./0.krakend-docker/)
- **Description:** Provides a minimal configuration file and instructions to run KrakenD using Docker.
- **Key Features:** Basic configuration, single endpoint setup, Docker commands.

### [1.playground](./1.playground/)
- **Description:** A complete demo environment built with Docker Compose, featuring multiple use-cases and integrations.
- **Key Features:** Demo setup, monitoring solutions, authentication integration, internal APIs for testing.

### [2.debug](./2.debug/)
- **Description:** Demonstrates how to use the KrakenD debug endpoint to get detailed request and response logs.
- **Key Features:** Debugging service, parameter forwarding, detailed logs.

### [3.flexible-configuration](./3.flexible-configuration/)
- **Description:** Example implementation of flexible configuration, including variables, templates, code snippets, and basic logic.
- **Key Features:** Docker setup, environment-based settings, immutable Docker artifact.

### [4.encodings](./4.encodings/)
- **Description:** Illustrates how to handle multiple response encodings and consolidate them using KrakenD.
- **Key Features:** Aggregating responses, encoding consolidation.

### [5.data-aggregation](./5.data-aggregation/)
- **Description:** Shows how to use KrakenD to aggregate data from multiple sources into a single endpoint.
- **Key Features:** Data aggregation, configuration examples.

### [6.traffic-throttling](./6.traffic-throttling/)
- **Description:** Demonstrates traffic throttling features of KrakenD, including rate limiting and circuit breaker patterns.
- **Key Features:** Rate limiting, circuit breakers.

### [7.backends-with-basic-auth](./7.backends-with-basic-auth/)
- **Description:** Example of connecting KrakenD with backends that require Basic Authentication.
- **Key Features:** Basic Authentication, backend configuration.

### [8.api-monetization-with-moesif](./8.api-monetization-with-moesif/)
- **Description:** Integration with Moesif for API analytics and monetization.
- **Key Features:** API event tracking, subscription management, usage-based billing.

### [9.rate-limits-per-tier](./9.rate-limits-per-tier/)
- **Description:** Proof of concept for custom rate limits based on user tiers using API keys and Lua scripting.
- **Key Features:** Custom rate limits, Lua scripting, API key authentication.

### [10.api-docs-with-redocly](./10.api-docs-with-redocly/)
- **Description:** Generates API documentation using Redocly and includes configuration examples.
- **Key Features:** API documentation generation.

### [lua_examples](./lua_examples/)
- **Description:** Various examples demonstrating Lua scripting in KrakenD for different use cases.
    - [lua_generate_header_from_host](./lua_examples/lua_generate_header_from_host/): Manipulate HTTP request headers dynamically using Lua scripting.
    - [lua_merge_grouby](./lua_examples/lua_merge_grouby/): Merge and group data from multiple sources.
    - [lua_soap_to_csv](./lua_examples/lua_soap_to_csv/): Convert SOAP responses to different formats (JSON, XML, CSV).
    - [lua_merge_and_paginate](./lua_examples/lua_merge_and_paginate/): Merge and paginate data from multiple sources.

### [multiple_post](./multiple_post/)
- **Description:** Examples demonstrating multiple sequential POST calls in KrakenD.
    - [enterprise](./multiple_post/enterprise/): Sequential POST calls using KrakenD Enterprise.
    - [community](./multiple_post/community/): Aggregating POST requests to SOAP backends.

### [plugins](./plugins/)
- **Description:** Examples demonstrating how to build and inject custom plugins into KrakenD.
    - [client](./plugins/client/): Custom HTTP client plugin.
    - [handler](./plugins/handler/): Custom HTTP handler plugin.
    - [modifier](./plugins/modifier/): Request/Response modifier plugin.

### [return_error_details](./return_error_details/)
- **Description:** Test environment for validating `X-KrakenD-Completed` headers and `return_error_details` configuration parameter.
- **Key Features:** Error details, incomplete response handling.

### [websockets](./websockets/)
- **Description:** Example integration of KrakenD with WebSocket protocol for real-time communication.
- **Key Features:** WebSocket configuration, real-time bi-directional communication.

---

Each example is designed to showcase specific features and use cases of KrakenD, providing you with practical implementations and configurations to help you get the most out of your API Gateway setup. For more details, refer to the individual README files linked above.
