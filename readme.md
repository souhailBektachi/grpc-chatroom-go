# gRPC with Go

This project demonstrates how to create a gRPC service using Go. It includes examples of defining a service, implementing the server, and creating a client to interact with the service.

## Prerequisites

- Protocol Buffers compiler (protoc)
- gRPC Go plugin for protoc

## Installation

1. Install Go from the [official website](https://golang.org/dl/).
2. Install the Protocol Buffers compiler from the [official website](https://developers.google.com/protocol-buffers).
3. Install the gRPC Go plugin:
    ```sh
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    ```

## Project Structure

```
grpcWithGo/
├── proto/
│   └── service.proto
├── server/
│   └── main.go
├── client/
│   └── main.go
└── readme.md
```

- `proto/service.proto`: Contains the service definition using Protocol Buffers.
- `server/main.go`: Implements the gRPC server.
- `client/main.go`: Implements the gRPC client.

## Usage


### Run the Server

Navigate to the `server` directory and run the server:

```sh
cd server
go run main.go
```

### Run the Client

Navigate to the `client` directory and run the client:

```sh
cd client
go run main.go
```
