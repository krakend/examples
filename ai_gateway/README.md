# KrakenD AI Gateway Examples

This repository contains practical examples demonstrating KrakenD's AI Gateway capabilities. Each example showcases different use cases and features for building production-ready AI gateways.

## Examples

### [LLM - Role-Based Multi-LLM Routing](./llm)

Demonstrates role-based routing to different AI providers (Gemini vs OpenAI) based on user authentication and JWT claims. This example shows how to build a unified API that routes requests to different LLM providers based on business logic.

**Key features**: JWT authentication, conditional backend routing, request/response templates, quota management per role

**Use case**: Organizations that need to route different user groups to different AI providers while maintaining a single API interface

---

### [RAG - Retrieval Augmented Generation](./rag)

Shows how to build a RAG server using KrakenD's sequential proxying capabilities. Integrates Weviate as a vector database and Gemini for both text embeddings and LLM responses.

**Key features**: Sequential backend calls, vector search integration, context-aware AI responses, request/response transformation

**Use case**: Applications that need to provide AI responses with context from a knowledge base or document repository

---

### [MCP - Model Context Protocol Server](./mcp)

Demonstrates KrakenD as an MCP server that aggregates multiple APIs (REST and GraphQL) to provide unified data access for AI agents. Combines country information from multiple sources with real-time weather data.

**Key features**: MCP protocol implementation, multi-API aggregation, data propagation between calls, JMESPath and Lua transformations

**Use case**: Providing AI agents with structured access to external tools and data sources through the MCP standard

---

### [LLM (Pre-2.11)](./llm.pre-2.11)

Legacy version of the role-based routing example for KrakenD versions before 2.11. Maintained for compatibility with older KrakenD installations.

**Key features**: Similar to the main LLM example but compatible with earlier KrakenD versions

**Use case**: Organizations running KrakenD versions prior to 2.11

## Common Prerequisites

All examples require:
- Docker and Docker Compose
- KrakenD Enterprise license (need a trial license? [Contact us](https://www.krakend.io/contact-sales/))

Some examples also require:
- API keys for AI providers (Gemini, OpenAI)
- Additional setup steps detailed in each example's README

## Getting Started

1. Choose an example that matches your use case
2. Navigate to the example directory
3. Follow the README instructions for that specific example
4. Each example is self-contained and can run independently

## Resources

- [KrakenD AI Gateway Documentation](https://www.krakend.io/docs/enterprise/ai-gateway/)
- [KrakenD Enterprise Documentation](https://www.krakend.io/docs/enterprise/)
- [Extended Flexible Configuration](https://www.krakend.io/docs/enterprise/configuration/flexible-config/)
