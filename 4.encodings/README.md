###### KrakenD - Quick Video Tutorials

# Working with multiple encodings in KrakenD

## Encoding consolidation on data aggregation

One of the main objectives of an API Gateway is help up to deal with complexity in our internal services. Eventually we must deal with APIs or services returning different formats and encodings, and our API Gateway can help us acting as a consolidation layer in front of the data consumers.

In this example you find a fake API that implements responses in multiple formats, running a LWAN server in a docker compose. KrakenD will expose two endpoints that will aggregate all that responses, one consolidating the answer in JSON and another consolidating in XML.

### Fake API Endpoints

| Response format | Endpoint                                                                           |
|-----------------|------------------------------------------------------------------------------------|
| JSON            | [`http://localhost:8000/response.json`](http://localhost:8000/response.json)       |
| XML             | [`http://localhost:8000/response.xml`](http://localhost:8000/response.xml)         |
| RSS             | [`http://localhost:8000/response-rss.xml`](http://localhost:8000/response-rss.xml) |
| Text            | [`http://localhost:8000/response.txt`](http://localhost:8000/response.txt)         |

### Consolidated KrakenD Endpoints

| Response format | Endpoint                                                                       |
|-----------------|--------------------------------------------------------------------------------|
| JSON            | [`http://localhost:8080/encodings.json`](http://localhost:8080/encodings.json) |
| XML             | [`http://localhost:8080/encodings.xml`](http://localhost:8080/encodings.xml)   |

We will combine the `enconding` parameter at backends scope, allowing KrakenD to know in what encoding each backend is answering, with the `ouput_encoding` paramater at endpoint scope, that will tell KrakenD in what enconding answer the requests.

_Note: you'll see that in our XML endpoint, to prevent repeating the full endpoint definition, we've just pointed our previous endpoint as a backend, modifying the `output_enconding`_

You can run this example by just executing a `docker-compose up` from the root folder.

You will find detailed information in our documentation:
- Output enconding: https://www.krakend.io/docs/endpoints/content-types/
- Supported backend encodings: https://www.krakend.io/docs/backends/supported-encodings/
