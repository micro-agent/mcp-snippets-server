# MCP Snippets Server

An MCP (Model Context Protocol) server that provides semantic search capabilities for code snippets using RAG (Retrieval-Augmented Generation) and vector embeddings.

## Overview

This server processes Markdown files containing code snippets, creates vector embeddings for semantic search, and provides an MCP tool for finding relevant code snippets based on topic queries.

## Features

- **Vector Store**: Creates and manages a persistent vector store from Markdown documentation
- **Semantic Search**: Uses OpenAI-compatible embeddings to find relevant snippets
- **MCP Integration**: Exposes search functionality as an MCP tool
- **Automatic Processing**: Processes `.md` files on first run and stores embeddings
- **Persistent Storage**: Saves vector store to JSON for quick subsequent startups

## Architecture

The server consists of several key components:

- **Vector Store** (`rag/`): Handles embedding storage and similarity search
- **Helpers** (`helpers/`): File processing utilities
- **Snippets** (`snippets/`): Contains code snippet documentation in Markdown format
- **MCP Server**: Exposes the `search_snippet` tool via HTTP

## Configuration

Set the following environment variables (or use `.env` file):

- `MODEL_RUNNER_BASE_URL`: OpenAI-compatible API endpoint (default: `http://localhost:12434/engines/llama.cpp/v1/`)
- `EMBEDDING_MODEL`: Embedding model name (default: `ai/mxbai-embed-large:latest`)
- `JSON_STORE_FILE_PATH`: Vector store file path (default: `rag-memory-store.json`)
- `MCP_HTTP_PORT`: HTTP server port (default: `9090`)
- `LIMIT`: Similarity threshold (default: `0.6`)
- `MAX_RESULTS`: Maximum search results (default: `2`)

## Usage

### Starting the Server

```bash
go run main.go
```

The server will:
1. Load existing vector store or create new one from `.md` files
2. Start HTTP server on the configured port
3. Expose MCP endpoint at `/mcp`

### MCP Tool

The server provides one MCP tool:

- **`search_snippet`**: Find code snippets related to a topic
  - Parameter: `topic` (string) - Search query or question

### Example Tool Call

```json
{
  "method": "tools/call",
  "params": {
    "name": "search_snippet",
    "arguments": {
      "topic": "how to create a REST API in Go"
    }
  }
}
```

## Development

### File Structure

- `main.go`: Main server implementation
- `rag/`: Vector store and similarity search logic
- `helpers/`: File processing utilities
- `snippets/`: Code snippet documentation
- `store/`: Persistent vector store data

### Adding New Snippets

1. Add Markdown files to the `snippets/` directory
2. Use `----------` as delimiter between different snippets
3. Restart the server to reprocess and update embeddings

## Dependencies

- `github.com/mark3labs/mcp-go`: MCP server implementation
- `github.com/openai/openai-go/v2`: OpenAI API client for embeddings
- `github.com/joho/godotenv`: Environment variable management
- `github.com/google/uuid`: UUID generation

## Docker Support

The project includes Docker configuration for easy deployment.

## Use the Docker Image

Image: https://hub.docker.com/repository/docker/k33g/mcp-snippets-server/tags

1. In a directory, create a `snippets` folder and add your `.md` files with code snippets.
2. Create a `compose.yml` file with the following content:

```yaml
services:

  mcp-snippets-server:
    image: k33g/mcp-snippets-server:0.0.1
    ports:
      - 9090:6060
    environment:
      - MCP_HTTP_PORT=6060
      - LIMIT=0.6
      - MAX_RESULTS=2
      - JSON_STORE_FILE_PATH=store/rag-memory-store.json
    volumes:
      - ./snippets:/app/snippets
      - ./store:/app/store

    models:
      mxbai-embed:
        endpoint_var: MODEL_RUNNER_BASE_URL
        model_var: EMBEDDING_MODEL

models:
  mxbai-embed:
    model: ai/mxbai-embed-large:latest
```

Start the server with:

```bash
docker compose up
```
