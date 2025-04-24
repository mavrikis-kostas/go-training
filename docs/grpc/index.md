# gRPC in Go

This section covers the basics of gRPC (Google Remote Procedure Call) in Go, with a focus on practical examples and best
practices for beginners.

## What You'll Learn

- **[Introduction to gRPC](introduction.md)**: Learn what gRPC is, its benefits, and when to use it.
- **[Protocol Buffers](protobuf.md)**: Understand Protocol Buffers (protobuf), the interface definition language used
  with gRPC.
- **[Server-Client Example](server-client-example.md)**: See a complete example of a gRPC server and client
  implementation in Go.
- **[Exercises](exercises.md)**: Practice what you've learned with hands-on exercises.

## Why gRPC Matters

gRPC is a modern, high-performance RPC (Remote Procedure Call) framework that can run in any environment. It enables
client and server applications to communicate transparently, making it easier to build connected systems.

Key benefits of gRPC include:

- **Efficient communication**: Uses HTTP/2 for transport and Protocol Buffers for serialization
- **Strong typing**: Interface contracts are defined using Protocol Buffers
- **Cross-platform**: Works across multiple languages and platforms
- **Bi-directional streaming**: Supports streaming in both directions

The examples in this section are designed to be simple and focused on gRPC concepts rather than complex application
logic, making them ideal for engineers who are new to gRPC.