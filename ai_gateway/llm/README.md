# KrakenD AI Gateway: Role-Based Multi-LLM Routing Demo

This demo showcases KrakenD's AI Gateway capabilities through a practical implementation of **role-based multi-LLM routing**. It demonstrates how to use KrakenD `ai/llm` feature as an unified interface to route requests to different AI providers (Gemini vs OpenAI) based on user authentication and roles.

An AI Gateway provides secure, scalable access to Large Language Models while abstracting vendor-specific APIs, controlling costs, and enforcing governance policies. This example shows how to implement these concepts in practice.

## What This Demonstrates

**Role-based AI routing**: Users with `position: developer` get routed to Gemini AI, while `position: support` users get OpenAI responses.

**Key capabilities**:
- JWT-based authentication and routing
- LLM provider selector using conditional backend
- Unified API for multiple AI providers  
- [Request](https://www.krakend.io/docs/enterprise/backends/body-generator/) and [response](https://www.krakend.io/docs/enterprise/backends/response-body-generator/) transformation using templates
- [Quota management](https://www.krakend.io/docs/enterprise/ai-gateway/budget-control/) per user type

## Quick Start

### Prerequisites

- Docker and Docker Compose
- KrakenD Enterprise license (need a trial license? [Contact us](https://www.krakend.io/contact-sales/))
- API keys for Gemini and OpenAI

### Setup

1. Create `config/krakend/.env`:
```bash
GEMINI_API_KEY=your_gemini_api_key
OPENAI_API_KEY=your_openai_api_key
```

2. Add your KrakenD Enterprise license as `LICENSE` in the root directory

3. Start services:
```bash
docker-compose up -d
```

**Important: It may take up to 2 minutes for all the services to properly start**

4. Open http://localhost:3000

### Test Users

**Developers → Gemini AI**:
- `tony` / `tony`
- `ana` / `ana` 
- `daisi` / `daisi`

**Support → OpenAI**:
- `holly` / `holly`
- `john` / `john`
- `matteo` / `matteo`

## How It Works

1. **Authentication**: Keycloak issues JWTs with user position claims
2. **Routing using conditional backends**: KrakenD reads the `position` claim and routes accordingly using [conditional backends](https://www.krakend.io/docs/enterprise/backends/conditional/#content):
   - `developer` → `POST` to Gemini API
   - `support` → `POST` to OpenAI API
3. **Templates**: [Request](https://www.krakend.io/docs/enterprise/backends/body-generator/) and [response](https://www.krakend.io/docs/enterprise/backends/response-body-generator/) bodies are transformed using Go templates
4. **Unified Response**: Both providers return the same JSON structure

## Configuration

The main routing logic is in `config/krakend/krakend.tmpl`. Templates for each provider are in `config/krakend/templates/vendors/`.

## Additional Features

**[Quota management](https://www.krakend.io/docs/enterprise/ai-gateway/budget-control/)**: Different token limits for admin vs regular users

**Monitoring**: Request/response logging and token usage tracking

## Resources

- [KrakenD AI Gateway Documentation](https://www.krakend.io/docs/enterprise/ai-gateway/)
- [Extended Flexible Configuration](https://www.krakend.io/docs/enterprise/configuration/flexible-config/)