# KrakenD AI Gateway: RAG server demo

This demo showcases a RAG server abstraction layer using KrakenD built-in capabilities. It uses Weviate as a VectorDB to store the documents and Gemini as the LLM and text embeddings provider.

**Key capabilities**:
- Sequential proxying and data propagation between backends
- [Request](https://www.krakend.io/docs/enterprise/backends/body-generator/) and [response](https://www.krakend.io/docs/enterprise/backends/response-body-generator/) transformation using templates

## Quick Start

### Prerequisites

- Docker and Docker Compose
- KrakenD Enterprise license
- API key for Gemini

### Setup

1. Create `config/krakend/.env`:
```bash
GEMINI_API_KEY=your_gemini_api_key
```

2. Add your KrakenD Enterprise license as `LICENSE` in the root directory

3. Start services:
```bash
docker-compose up -d
```

4. Add documents and query with context using `/rag/add` and `/rag/query`

**Important: It may take up to a minute for all the services to properly start**

## How It Works

### Adding documents
1. The input documents are transformed into vector info using Gemini text embedding API
2. The documents and their vector info is sent to weviate

**cURL example**
``` 
$ curl -H'Content-type: application/json' http://localhost:8080/rag/add -d '{"documents":[{"text": "the x043 diturium motors can yield up to 14500 hexagonal rifts"}, {"text": "do not operate a x043 motor transfussor above 435 anglar degrees"}]}'

{"added":2}
```

### Querying
1. The input query is transformed into vector info using Gemini text embedding API
2. The vector info is used to perform a near search using weviate. If a match is found, the documents are returned
3. A prompt is generated using templates and the documents from previous step
4. The prompt is sent to Gemini and the results are transformed into a unified response using templates

**cURL example**
```
$ curl -H'Content-type: application/json' http://localhost:8080/rag/query -d '{"content": "can a x043 diturium motor work at 500 degrees?"}'

{"ai_gateway_response":[{"contents":["No, do not operate a x043 motor transfussor above 435 anglar degrees.\n"]}],"usage":"212"}
``` 

## Configuration

The main routing logic is in `config/krakend/krakend.tmpl`. Templates for each provider are in `config/krakend/templates/vendors/`.

## Additional Features

**Monitoring**: Request/response logging and token usage tracking

## Resources

- [KrakenD AI Gateway Documentation](https://www.krakend.io/docs/enterprise/ai-gateway/)
- [Extended Flexible Configuration](https://www.krakend.io/docs/enterprise/configuration/flexible-config/)
- [Sequential proxying](https://www.krakend.io/docs/endpoints/sequential-proxy/)