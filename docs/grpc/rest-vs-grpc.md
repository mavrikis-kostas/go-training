## REST vs. gRPC: A Simple `SayHello` Example

### Overview of Key Differences

| Feature         | REST                            | gRPC                                    |
|-----------------|---------------------------------|-----------------------------------------|
| Transport       | HTTP 1.1                        | HTTP/2 (supports multiplexing, streams) |
| Message Format  | JSON                            | Protobuf (binary, smaller & faster)     |
| API Definition  | No strict schema                | Strongly typed via `.proto` files       |
| Code Generation | Manual request/response structs | Auto-generated from `.proto`            |
| Performance     | Text-based (slower, larger)     | Binary (fast, compact)                  |
| Tooling         | curl, Postman                   | grpcurl, evans, generated clients       |

---

## The gRPC Example (`SayHello`)

### `hello.proto`

```protobuf
syntax = "proto3";

package hello;

option go_package = "example.com/hello";

service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply);
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 1;
}
```

Generate the code:

```bash
protoc --go_out=. --go-grpc_out=. hello.proto
```

### gRPC Server (`grpc_server.go`)

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "example.com/hello"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %s", req.Name)
	return &pb.HelloReply{Message: "Hello " + req.Name}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Println("gRPC Server running on :50051")
	s.Serve(lis)
}
```

---

### gRPC Client (`grpc_client.go`)

```go
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "example.com/hello"
)

func main() {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, _ := client.SayHello(ctx, &pb.HelloRequest{Name: "Alice"})
	log.Printf("Response: %s", resp.Message)
}
```

## REST Example (`SayHello`)

### Server (`rest_server.go`)

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	var req HelloRequest
	json.NewDecoder(r.Body).Decode(&req)
	res := HelloResponse{Message: "Hello " + req.Name}
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/sayhello", sayHello)
	log.Println("REST Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
```

### REST Client (`rest_client.go`)

```go
package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	reqBody := map[string]string{"name": "Alice"}
	jsonData, _ := json.Marshal(reqBody)

	resp, err := http.Post("http://localhost:8080/sayhello", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var res map[string]string
	json.NewDecoder(resp.Body).Decode(&res)
	log.Printf("Response: %s", res["message"])
}
```

## Summary: When to Use What

- **gRPC**: Internal services, low-latency communication, strong typing, automatic codegen, microservices.
- **REST**: Public APIs, browser-based clients, wide tooling support, simpler onboarding.
