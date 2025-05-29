# Simple API Server

A lightweight HTTP server written in Go that provides an API endpoint for receiving and logging messages from the Jarvis agent. This server is designed to handle log messages and provide a simple interface for testing and development purposes.

## Features

- HTTP server running on port 8080
- RESTful API endpoint for log messages
- JSON request/response handling
- Dual logging (file and console output)
- Support for both POST and GET methods

## API Endpoints

### POST /api/v1/logs
Receives and logs messages from the Jarvis agent.

**Request Body Format:**
```json
{
    "timestamp": "2024-03-21T10:00:00Z",
    "source": "jarvis-agent-windows",
    "type": "log",
    "level": "info",
    "message": "Example message",
    "data": {},
    "agent_id": "agent-123",
    "environment": "development"
}
```

**Response:**
```json
{
    "status": "success",
    "message": "Log received"
}
```

### GET /api/v1/logs
Currently returns a simple status message (retrieval functionality not implemented).

## Prerequisites

- Go 1.23.1 or later
- Basic understanding of HTTP APIs and JSON

## Building and Running

1. Clone the repository:
```bash
git clone <repository-url>
cd SimpleApiServer
```

2. Build the application:
```bash
go build
```

3. Run the server:
```bash
./SimpleApiServer
```

The server will start on port 8080 and begin logging to both the console and `api_server.log` file.

## Logging

The server maintains two types of logs:
- Console output: Shows real-time logs in the terminal
- File output: Writes logs to `api_server.log` in the application directory

## Error Handling

The server includes basic error handling for:
- Invalid JSON payloads
- Unsupported HTTP methods
- File system errors
- Server startup issues

## Development

To modify or extend the server:
1. Edit `main.go` to add new endpoints or functionality
2. Update the `LogMessage` struct if the message format changes
3. Rebuild and restart the server

## License

[Add your license information here] 