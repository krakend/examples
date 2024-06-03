## Implementing Multiple Sequential POST Calls with KrakenD

This example demonstrates how to use KrakenD to implement multiple POST calls in a sequential proxy, dynamically transform data from a GET request into a POST call, inject data from one service's response into another POST request, and aggregate the responses.

### Use Case

We have two services:

1. **User Service**: Returns user data based on a `username`.
2. **Reviews Service**: Returns all reviews from a user based on their `user_id`.

We'll implement a single GET endpoint in KrakenD that:

1. Extracts the `username` from the GET request path.
2. Uses this `username` to make a POST request to the User Service.
3. Takes the `user_id` from the User Service response and uses it to make a POST request to the Reviews Service.
4. Aggregates the responses from both services into a single response.

### Prerequisites

- KrakenD Enterprise license (required for dynamic routing and request modifiers with templates).
- Docker and Docker Compose installed.

### Files in the Repository

- `docker-compose.yml`: Docker Compose file including KrakenD Enterprise and basic services to simulate the User and Reviews services.
- `index.php`: Basic implementation of the User and Reviews services.
- `krakend.json`: KrakenD configuration file.
- `LICENSE`: Placeholder for a valid KrakenD Enterprise license.
- `reviews_post_body.tmpl`: Template for the Reviews Service POST request body.
- `users_post_body.tmpl`: Template for the User Service POST request body.

### Running the Example

1. **Start the Services**:

    ```bash
    docker-compose up
    ```

2. **Make a GET Request**:

    ```bash
    curl http://localhost:8080/user-and-reviews/admin
    ```


### Expected Behavior

1. The GET request to `/user-and-reviews/{username}` will extract the `username` from the path.
2. KrakenD will use this `username` to make a POST request to the User Service.
3. The `user_id` from the User Service response will be extracted and used to make a POST request to the Reviews Service.
4. The responses from both services will be aggregated and returned in a single response.

### Further Reading

- [Sequential Proxy Documentation](https://www.krakend.io/docs/enterprise/endpoints/sequential-proxy/)
- [Request Modifier with Templates Documentation](https://www.krakend.io/docs/enterprise/backends/body-generator/)

### Notes

- Ensure you replace the `LICENSE` file with a valid KrakenD Enterprise license to run this example.
- Adjust the paths and configuration in `krakend.json` as needed to match your setup.

This example showcases KrakenD's powerful capabilities to handle complex API orchestration scenarios with ease. For more details, visit [krakend.io](https://www.krakend.io/).
