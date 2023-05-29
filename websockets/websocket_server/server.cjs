const WebSocket = require('ws');

const wss = new WebSocket.Server({ port: 5678 });

wss.on('connection', ws => {
    ws.on('message', message => {
        console.log("\n\nNew Request: %s", message);
        try{
            var json_message = JSON.parse(message);
            if (json_message.body)
            {
                let response = {};
                response.body = json_message.body;

                if (json_message.session.uuid)
                    response.session = { uuid: json_message.session.uuid };

                if (json_message.url)
                    response.url = json_message.url;

                let buff = Buffer.from(json_message.body, 'base64');
                let string_body = buff.toString('ascii');
                console.log('Message body: %s', string_body);

                let json_answer = JSON.stringify(response);

                console.log('Answer: %s', json_answer);
                ws.send(json_answer);
            }
        }catch(e){
            console.log('Error managing: %s', message);
        }

    });

    ws.send('OK');
});

console.log('WebSocket server running on port 5678');
