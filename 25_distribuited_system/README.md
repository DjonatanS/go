# Distributed System

A simple distributed system implementation with a master-worker architecture using gRPC for communication between nodes.

## Project Overview

This project implements a basic distributed system with two types of nodes:

1. **Master Node**: Coordinates tasks and distributes them to worker nodes
2. **Worker Node**: Executes tasks received from the master node

The system uses gRPC for communication between nodes and provides a REST API to submit tasks to the master.

## Architecture

![Architecture Diagram](https://mermaid.ink/img/pako:eNp1kMFqwzAMhl9F6JRC3sB0kFMHO5WVHbLTUFAsxRHzLWPJgZLk3ac4pRt0O-n_vk-WpBtYNgKdcs8DNx1rXggjaT-Q8ba1IgO50HAEGPAPtLbDZxAJ1eMMUhdC8MMID3U98bACvTAeDY2SxnraGB-6e1zHw-cY7oQVSTtBvlELWn9DjcFigJM0FCXlKFGK0PyiqOICXZg4FNkdlmX5IbYsd4VYaZW6c6nD8j_cYns-HS6fh8v35pTfbWTy3CyEpXGkW3RENB3-_TQO5o0pt7a_SLdsOno1tWLkocQbdFzdi9RT8hSA83jpTQXBGxVC14VOk8uSrBedtKLEpMNXl5w2PocQXTb8BT5OgHs)

### Components

- **Master Node**: 
  - Listens for gRPC connections from worker nodes
  - Provides an HTTP API for task submission
  - Distributes tasks to connected worker nodes

- **Worker Node**: 
  - Connects to the master via gRPC
  - Receives and executes tasks (commands)
  - Reports status back to the master

## Getting Started

### Prerequisites

- Go 1.18 or later
- protoc (Protocol Buffers compiler)

### Installation

```bash
# Clone the repository
git clone <repository-url>
cd 25_distribuited_system

# Install dependencies
go mod tidy
```

## Usage

The system can be started in either master or worker mode:

### Start the Master Node

```bash
go run main.go master
```

This will start:
- A gRPC server on port 50051 for worker connections
- A REST API on port 9092 for task submission

### Start a Worker Node

```bash
go run main.go worker
```

This will start a worker that connects to the master node at localhost:50051.

### Submit Tasks

You can submit tasks to the master node using the REST API:

```bash
# Submit a command to be executed by workers
curl -X POST http://localhost:9092/tasks \
  -H "Content-Type: application/json" \
  -d '{"cmd":"echo hello world"}'
```

## API Reference

### REST API

| Endpoint | Method | Request Body | Description |
|----------|--------|--------------|-------------|
| `/tasks` | POST   | `{"cmd":"command_to_execute"}` | Submit a task to be executed by worker nodes |

### gRPC Services

The system defines the following gRPC services:

```protobuf
service NodeService {
    rpc ReportStatus(Request) returns (Response){};
    rpc AssignTask(Request) returns (stream Response){};
}
```

## How It Works

1. **Initialization**:
   - The master node starts and listens for gRPC connections
   - Worker nodes connect to the master via gRPC

2. **Task Assignment**:
   - Tasks are submitted via HTTP to the master's REST API
   - The master sends tasks to connected workers through the `AssignTask` gRPC stream

3. **Command Execution**:
   - Workers receive commands and execute them using the system's shell
   - The execution output is printed on the worker's console

## Project Structure

```
.
├── core/
│   ├── node.go          # Contains master node and service implementations
│   ├── node.pb.go       # Generated protobuf message definitions
│   ├── node_grpc.pb.go  # Generated gRPC service definitions
│   ├── node.proto       # Protocol buffer definitions
│   └── worker_node.go   # Worker node implementation
├── go.mod
├── go.sum
└── main.go              # Entry point
```

## Technical Details

### Master Node Implementation

The master node:
- Creates a gRPC server for worker connections
- Maintains a channel for distributing commands
- Provides a REST API for task submission (using Gin)
- Broadcasts commands to all connected workers

### Worker Node Implementation

The worker node:
- Establishes a gRPC connection to the master
- Creates a streaming connection for receiving tasks
- Executes commands using Go's `exec` package
- Reports status to the master (currently minimal implementation)

### Communication Flow

1. Client submits a task to master's REST API
2. Master forwards the task to all connected workers via gRPC streams
3. Workers execute the commands and report status

## Development

### Regenerating Protocol Buffers

If you modify the `.proto` files, regenerate the Go code:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    core/node.proto
```

## Limitations and Future Improvements

- Currently, the same task is sent to all workers (no load balancing)
- Tasks are simple shell commands with no advanced scheduling
- No authentication or security measures
- Limited error handling and retry mechanisms
- No persistence of tasks or results

## License

[License information]
