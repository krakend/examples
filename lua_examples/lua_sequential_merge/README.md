# KrakenD Lua Scripting Example

This repository demonstrates how to use **KrakenD** to aggregate data from multiple microservices and manipulate the response using **Lua scripting**. It includes an example setup using **Docker Compose** with a mocked API and a KrakenD gateway, configured to merge product and review data into a single response.

## Architecture

The setup consists of:

- **KrakenD Gateway**: The entry point to aggregate and manipulate responses from multiple services.
- **Mocked Services**: Simple services implemented using BusyBox serving static JSON files.

## Services

### Product Service

Serves product information including name, price, and images.

Example Response:

```json
{
  "id": "8d5f5b10-6f34-4421-a9ee-7a3a0e5a5be8",
  "name": "Macbook Pro",
  "price": {
    "amount": "19.99",
    "currency": "USD"
  },
  "imageUrls": [
    "<https://example.com/image1.jpg>",
    "<https://example.com/image2.jpg>"
  ]
}
```

### Reviews Service

Serves reviews related to the products.

Example Response:

```json
[
  {
    "id": "1f9ead6d-e3eb-4f08-98a6-f30ab41a36f4",
    "userId": "f4555bb0-c743-44f6-92a3-376e0f90df06",
    "productId": "8d5f5b10-6f34-4421-a9ee-7a3a0e5a5be8",
    "rating": 4.5,
    "comment": "It's amazing"
  },
  {
    "id": "3a0591e2-e762-4427-97d8-65f6fc7cc1fc",
    "userId": "3e35e2c6-f58b-4752-a41f-9b77bc0b7a02",
    "productId": "8d5f5b10-6f34-4421-a9ee-7a3a0e5a5be8",
    "rating": 3.5,
    "comment": "Not the best I've used."
  },
  {
    "id": "c43e5a07-3add-40b3-8b46-8b6454938a72",
    "userId": "7e685cd9-1195-4bb8-ab10-02d4c5c36960",
    "productId": "8d5f5b10-6f34-4421-a9ee-7a3a0e5a5be8",
    "rating": 5,
    "comment": "Lovely!"
  }
]
```

### Users Service

Serves user data based on the `userId` present in the reviews.

Example Response:

```json
[
  {
    "id": "f4555bb0-c743-44f6-92a3-376e0f90df06",
    "name": "John Doe",
    "email": "john@doe.com",
    "profilePicture": "https://example.com/profile1.jpg"
  },
  {
    "id": "3e35e2c6-f58b-4752-a41f-9b77bc0b7a02",
    "name": "Unknown Guy",
    "email": "guy@unknown.com",
    "profilePicture": "https://example.com/profile2.jpg"
  },
  {
    "id": "7e685cd9-1195-4bb8-ab10-02d4c5c36960",
    "name": "Molly",
    "email": "molly@email.com",
    "profilePicture": "https://example.com/profile3.jpg"
  }
]
```

## Lua Script

The Lua script `merge_reviewers.lua` enhances the reviews by appending the reviewer details fetched from the **Users Service**.

### Functionality:

1. Extracts user IDs from the reviews.
2. Fetches user details from the **Users Service**.
3. Merges the user details (name, email, and profile picture) into the corresponding review.

Example Lua script:

```lua
function merge_reviewers(resp)local response_body = resp:data()
    local reviews = response_body:get("reviews")
    local user_ids = {}

    for i = 1, reviews:len() do
        local review = reviews:get(i - 1)
        table.insert(user_ids, review:get("userId"))
    end
    local user_ids_str = table.concat(user_ids, ",")
    local user_response = http_response.new("http://my_service/users.json?ids=" .. user_ids_str)

    if user_response:statusCode() == 200 then
        local user_data_str = user_response:body()
        local user_data = json_parse(user_data_str)

        for i = 1, reviews:len() do
            local review = reviews:get(i - 1)
            local user_id = review:get("userId")
            for _, user in ipairs(user_data) do
                if user.id == user_id then
                    review:set("userName", user.name)
                    review:set("userEmail", user.email)
                    review:set("userProfilePicture", user.profilePicture)
                end
            end
        end
    else
        print("Failed to fetch user data")
    end
end
```

## Final Aggregated Response

KrakenD will aggregate the responses from the **Product Service** and **Reviews Service**, enrich the reviews with user details from the **Users Service**, and return the following response:

```json
{
  "id": "8d5f5b10-6f34-4421-a9ee-7a3a0e5a5be8",
  "name": "Macbook Pro",
  "price": {
    "amount": "19.99",
    "currency": "USD"
  },
  "imageUrls": [
    "https://example.com/image1.jpg",
    "https://example.com/image2.jpg"
  ],
  "reviews": [
    {
      "id": "1f9ead6d-e3eb-4f08-98a6-f30ab41a36f4",
      "userId": "f4555bb0-c743-44f6-92a3-376e0f90df06",
      "productId": "8d5f5b10-6f34-4421-a9ee-7a3a0e5a5be8",
      "rating": 4.5,
      "comment": "It's amazing",
      "userName": "John Doe",
      "userEmail": "john@doe.com",
      "userProfilePicture": "https://example.com/profile1.jpg"
    },
    {
      "id": "3a0591e2-e762-4427-97d8-65f6fc7cc1fc",
      "userId": "3e35e2c6-f58b-4752-a41f-9b77bc0b7a02",
      "productId": "8d5f5b10-6f34-4421-a9ee-7a3a0e5a5be8",
      "rating": 3.5,
      "comment": "Not the best I've used.",
      "userName": "Unknown Guy",
      "userEmail": "guy@unknown.com",
      "userProfilePicture": "https://example.com/profile2.jpg"
    },
    {
      "id": "c43e5a07-3add-40b3-8b46-8b6454938a72",
      "userId": "7e685cd9-1195-4bb8-ab10-02d4c5c36960",
      "productId": "8d5f5b10-6f34-4421-a9ee-7a3a0e5a5be8",
      "rating": 5,
      "comment": "Lovely!",
      "userName": "Molly",
      "userEmail": "molly@email.com",
      "userProfilePicture": "https://example.com/profile3.jpg"
    }
  ]
}
```

## Running the Project

1. **Clone the repository**:

    ```bash
    git clone https://github.com/krakend/examples.git
    cd lua_examples/lua_sequential_merge
    ```

2. **Run the services**:
   Use Docker Compose to bring up the services:

    ```bash
    docker-compose up
    ```

3. **Test the Gateway**:
   Once the services are running, you can test the aggregated response:

    ```bash
    curl http://localhost:8080/full-product-data/8d5f5b10-6f34-4421-a9ee-7a3a0e5a5be8
    ```


You should see the fully aggregated and enriched product data with reviews and user information.
