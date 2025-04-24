# gRPC Server and Client Example

This guide walks through building a simple gRPC server and client in Go. We'll implement the Calculator service defined
in the [Protocol Buffers](protobuf.md) section.

## Project Structure

First, let's set up our project structure:

```
calculator/
├── calculator.proto	 # Protocol Buffers definition
├── server/
│	└── main.go			# gRPC server implementation
├── client/
│	└── main.go			# gRPC client implementation
└── example.com/calculator/			# Generated code (will be created)
```

## Step 1: Define the Service with Protocol Buffers

Create a file named `calculator.proto`:

```protobuf
syntax = "proto3";

package calculator;

option go_package = "example.com/calculator";

// Calculator service definition
service Calculator {
  // Unary RPC
  rpc Add (AddRequest) returns (AddResponse);

  // Unary RPC
  rpc Subtract (SubtractRequest) returns (SubtractResponse);
}

// Request message for Add method
message AddRequest {
  int32 a = 1;
  int32 b = 2;
}

// Response message for Add method
message AddResponse {
  int32 result = 1;
  string operation = 2;  // Description of the operation performed
}

// Request message for Subtract method
message SubtractRequest {
  int32 a = 1;
  int32 b = 2;
}

// Response message for Subtract method
message SubtractResponse {
  int32 result = 1;
  string operation = 2;  // Description of the operation performed
}
```

## Step 2: Generate Go Code from the .proto File

Run the following command to generate Go code from your `.proto` file:

```bash
protoc --go_out=. --go-grpc_out=. calculator.proto
```

This will create a `calculator` directory with two files:

- `calculator.pb.go`: Contains Go structs for your messages
- `calculator_grpc.pb.go`: Contains client and server interfaces for your service

## Step 3: Implement the gRPC Server

Create a file named `server/main.go`:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "go-together/example.com/calculator"
)

// server is used to implement the Calculator service
type server struct {
	pb.UnimplementedCalculatorServer
}

// Add implements the Add RPC method
func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("Received: %v + %v", req.A, req.B)
	return &pb.AddResponse{
		Result:    req.A + req.B,
		Operation: fmt.Sprintf("Added %d and %d", req.A, req.B),
	}, nil
}

// Subtract implements the Subtract RPC method
func (s *server) Subtract(ctx context.Context, req *pb.SubtractRequest) (*pb.SubtractResponse, error) {
	log.Printf("Received: %v - %v", req.A, req.B)
	return &pb.SubtractResponse{
		Result:    req.A - req.B,
		Operation: fmt.Sprintf("Subtracted %d from %d", req.B, req.A),
	}, nil
}

func main() {
	// Create a TCP listener on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register our service with the gRPC server
	pb.RegisterCalculatorServer(s, &server{})

	// Start serving requests
	fmt.Println("Server started on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
```

## Step 4: Implement the gRPC Client

Create a file named `client/main.go`:

```go
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "go-together/example.com/calculator"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := pb.NewCalculatorClient(conn)

	// Set a timeout for our API call
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the Add method
	addResponse, err := client.Add(ctx, &pb.AddRequest{A: 10, B: 20})
	if err != nil {
		log.Fatalf("Add failed: %v", err)
	}
	log.Printf("Add result: %v (%s)", addResponse.Result, addResponse.Operation)

	// Call the Subtract method
	subtractResponse, err := client.Subtract(ctx, &pb.SubtractRequest{A: 30, B: 5})
	if err != nil {
		log.Fatalf("Subtract failed: %v", err)
	}
	log.Printf("Subtract result: %v (%s)", subtractResponse.Result, subtractResponse.Operation)
}
```

## Step 5: Run the Example

1. Start the server:
   ```bash
   go run server/main.go
   ```

2. In another terminal, run the client:
   ```bash
   go run client/main.go
   ```

You should see output similar to:

**Server Output:**

```
Server started on :50051
Received: 10 + 20
Received: 30 - 5
```

**Client Output:**

```
Add result: 30 (Added 10 and 20)
Subtract result: 25 (Subtracted 5 from 30)
```

## Understanding the Example

### Server Implementation

1. We define a `server` struct that implements the `CalculatorServer` interface.
2. We implement the `Add` and `Subtract` methods, which handle incoming requests.
3. We create a TCP listener on port 50051 and register our service with the gRPC server.
4. We start the server to listen for incoming requests.
5. We log the received requests and return the results in the response messages.

### Client Implementation

1. We create a connection to the server using `grpc.NewClient`.
2. We create a client for the `Calculator` service.
3. We call the `Add` and `Subtract` methods on the client, passing in the request messages.

### Context and Timeout

Go's `context` package is used to manage timeouts and cancellation signals. In this example, we set a timeout of 1
second for our API calls using `context.WithTimeout`. This ensures that if the server does not respond within
the specified time, the client will cancel the request and return an error.

The `ctx` variable needs to be passed to the gRPC methods to ensure that the server can handle cancellation signals and
timeouts properly.

For more information on using context in Go, refer to the [context package documentation](https://pkg.go.dev/context).
