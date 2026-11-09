// Minimal Socket.IO chat server for the KrakenD example.
// It serves the chat page at "/" and broadcasts every "chat message" event.
const { createServer } = require("node:http");
const { readFileSync } = require("node:fs");
const { Server } = require("socket.io");

const page = readFileSync(__dirname + "/index.html");

const server = createServer((req, res) => {
    if (req.url === "/") {
        res.writeHead(200, { "Content-Type": "text/html" });
        res.end(page);
        return;
    }
    res.writeHead(404);
    res.end();
});

const io = new Server(server);
let serverOffset = 0;

io.on("connection", (socket) => {
    console.log("client connected:", socket.id);
    socket.on("chat message", (msg) => {
        serverOffset++;
        io.emit("chat message", msg, serverOffset);
    });
    socket.on("disconnect", () => console.log("client disconnected:", socket.id));
});

server.listen(3000, () => console.log("Socket.IO server running on port 3000"));
