# gRPC Server and Client Example

This guide walks through building a simple gRPC server and client in Go. We'll implement the Calculator service defined
in the [Protocol Buffers](protobuf.md) section.

## Project Structure

First, let's set up our project structure:

```
calculator/
├── calculator.proto    # Protocol Buffers definition
├── server/
│   └── main.go         # gRPC server implementation
├── client/
│   └── main.go         # gRPC client implementation
└── calculator/         # Generated code (will be created)
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

  // Server streaming RPC
  rpc CountDown (CountDownRequest) returns (stream CountDownResponse);
}

// Request message for Add method
message AddRequest {
  int32 a = 1;
  int32 b = 2;
}

// Response message for Add method
message AddResponse {
  int32 result = 1;
}

// Request message for CountDown method
message CountDownRequest {
  int32 start = 1;  // Start counting down from this number
}

// Response message for CountDown method
message CountDownResponse {
  int32 number = 1;  // Current number in the countdown
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
	"time"

	"google.golang.org/grpc"

	pb "example.com/calculator/calculator"
)

// server is used to implement the Calculator service
type server struct {
	pb.UnimplementedCalculatorServer
}

// Add implements the Add RPC method
func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("Received: %v + %v", req.A, req.B)
	return &pb.AddResponse{Result: req.A + req.B}, nil
}

// CountDown implements the CountDown RPC method
func (s *server) CountDown(req *pb.CountDownRequest, stream pb.Calculator_CountDownServer) error {
	log.Printf("Starting countdown from %v", req.Start)

	// Simple countdown
	for num := req.Start; num > 0; num-- {
		// Send the current number to the client
		if err := stream.Send(&pb.CountDownResponse{Number: num}); err != nil {
			return err
		}
		// Small delay to simulate processing time
		time.Sleep(100 * time.Millisecond)
	}
	return nil
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
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "example.com/calculator/calculator"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
	log.Printf("Add result: %v", addResponse.Result)

	// Call the CountDown method
	countDownStream, err := client.CountDown(ctx, &pb.CountDownRequest{Start: 10})
	if err != nil {
		log.Fatalf("CountDown failed: %v", err)
	}

	// Receive and print all numbers in the countdown
	log.Println("Countdown:")
	for {
		countDown, err := countDownStream.Recv()
		if err == io.EOF {
			// End of stream
			break
		}
		if err != nil {
			log.Fatalf("Error while receiving: %v", err)
		}
		log.Printf("  %v", countDown.Number)
	}
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

**Server output:**

```
Server started on :50051
Received: 10 + 20
Starting countdown from 10
```

**Client output:**

```
Add result: 30
Countdown:
  10
  9
  8
  7
  6
  5
  4
  3
  2
  1
```

## Understanding the Example

### Server Implementation

1. We create a gRPC server and register our Calculator service.
2. We implement the `Add` method for unary RPC (simple request-response).
3. We implement the `CountDown` method for server streaming RPC, where the server sends multiple responses.

### Client Implementation

1. We establish a connection to the gRPC server.
2. We create a client stub for the Calculator service.
3. We call the `Add` method and receive a single response.
4. We call the `CountDown` method and receive a stream of responses.

## Next Steps

This example demonstrates the basics of gRPC in Go. To build on this knowledge, you might want to explore:

1. **Client Streaming**: Where the client sends multiple requests to the server.
2. **Bidirectional Streaming**: Where both client and server send multiple messages.
3. **Error Handling**: How to properly handle and propagate errors in gRPC.
4. **Authentication**: How to secure your gRPC services.
5. **Interceptors**: Similar to middleware in HTTP servers, for cross-cutting concerns.

In the next section, we'll provide exercises to help you practice these concepts.