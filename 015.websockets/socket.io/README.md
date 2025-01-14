# **KrakenD Example with Socket.IO**

This repository provides an example of how to integrate KrakenD with Socket.IO. 
Please see the documentation here: https://www.krakend.io/docs/enterprise/websockets/#integrating-krakend-with-socketio

> [Socket.IO](https://socket.io/) is a popular library to use bidirectional communication. Although Socket.IO name might sound as a WebSockets implementation, the reality is that Socket.IO operates on a **custom protocol** layered over WebSockets that is incompatible with plain WebSockets clients using the [WebSockets API](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API/Writing_WebSocket_client_applications) (the one native in the JS standard lib). To connect to a Sockets.IO server you cannot use a WebSockets client, you must use a Sockets.IO client.

> KrakenD uses a pure WebSocket Protocol (RFC-6455) to connect to servers, but the Socket.IO protocol requires specific signaling to establish and maintain connections. By default, it attempts to **use the same endpoint for both HTTP and WebSocket communication**, and the connection details passed on a query string (e.g., `?EIO=4&transport=websocket`). This design can cause confusion when integrating with KrakenD, which manages HTTP and WebSocket traffic separately. Make sure to use `websockets` only when passing through KrakenD.

> Socket.IO also requires **dedicated connections for each client**. This approach is **incompatible with KrakenD's multiplexing system**, which optimizes resource usage by sharing WebSocket connections among multiple clients, so you are limited to use **direct WebSockets only**. Needles to say that handling individual client connections, leads to a much higher resource consumption.

> Integrating KrakenD with Socket.IO can open up powerful real-time communication features, but it comes with trade-offs. The need for dedicated per-client connections, the additional dependency footprint, and challenges in maintaining asynchronous logic and multi-threaded execution must be considered before committing to this setup.

