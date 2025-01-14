# **KrakenD WebSocket Integration Example**

This repository provides an example of how to integrate KrakenD, an ultra performant API Gateway, with WebSocket protocol. The configuration is designed to facilitate real-time bi-directional communication between the client and server.

## **Structure**

The repository includes a KrakenD configuration file (**`krakend.json`**), a **`docker-compose.yml`** file for setting up the infrastructure, and a WebSocket server implementation (**`/websocket_server/server.cjs`**).

## **KrakenD Configuration**

The **`krakend.json`** is a configuration file for KrakenD. It specifies gateway settings such as port, timeout, caching, debugging, and more. In this example, a WebSocket endpoint (**`/ws`**) is set up. This endpoint redirects to a WebSocket service at **`ws://ws_service:5678`**.

## **Docker Compose**

The **`docker-compose.yml`** file sets up two services: **`krakend`** and **`ws_service`**.

- **`krakend`** service runs the KrakenD gateway using the **`krakend/krakend-ee:watch`** image. It mounts the current directory to **`/etc/krakend`** inside the container and exposes port **`8080`**.
- **`ws_service`** is a custom WebSocket service, implemented using the Node.js **`ws`** library. It runs on port **`5678`** and is defined in **`/websocket_server/server.cjs`**.

## **WebSocket Server Implementation**

The **`/websocket_server/server.cjs`** is an illustrative WebSocket server implementation using the Node.js **`ws`** library. This server is part of the **`ws_service`** defined in the **`docker-compose.yml`** file and runs automatically when the Docker services are started.

KrakenD has specific requirements regarding message formatting due to its multiplexing feature. All messages between KrakenD and the WebSocket backend should include some envelopes, which allows for deciding whether the message should be sent to a specific user or broadcast to multiple clients. For more information, refer to the **[KrakenD WebSocket documentation](https://www.krakend.io/docs/enterprise/websockets/#message-format)**.

## **Prerequisites**

- Docker and Docker Compose installed on your machine
- A valid KrakenD Enterprise LICENSE file. You can [contact us to ask for a trial](https://www.krakend.io/enterprise/#contact-sales).

## **Usage**

To use this repository, follow these steps:

1. Clone the repository to your local machine.
2. Navigate to the project directory.
3. Run the docker compose command to start the services.

```
bash
$ git clone https://github.com/krakend/examples
$ cd websockets
$ docker-compose up

```

Once the services are up and running, you can access the WebSocket service through KrakenD API Gateway using **`ws://localhost:8080/ws`**.

You can use https://websocketking.com/ or any other WebSocket client to test this example. This test server will just answer with the same body received in the request.

## **Support**

If you encounter any problems or have any suggestions, please open an issue on this repository.
