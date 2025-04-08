# Introduction to net/http

Go's standard library includes a powerful package called `net/http` that provides HTTP client and server implementations. This package makes it easy to send HTTP requests and handle HTTP responses in your Go applications.

## Key Components of net/http

### The Client

The `http.Client` type is used to make HTTP requests. It manages connections, cookies, redirects, and other aspects of HTTP communication.

Here's how you can create and configure an HTTP client:

```go
// Creating a default HTTP client
client := &http.Client{}

// Creating a client with custom timeout
client := &http.Client{
    Timeout: 10 * time.Second,
}
```

### The Request

The `http.Request` type represents an HTTP request to be sent by a client or received by a server.

Here's how you can create and configure an HTTP request:

```go
// Creating a new HTTP request
req, err := http.NewRequest("GET", "https://api.example.com/data", nil)
if err != nil {
    // Handle error
}

// Adding headers to the request
req.Header.Add("Content-Type", "application/json")
req.Header.Add("Authorization", "Bearer token123")
```

### The Response

The `http.Response` type represents the response from an HTTP request.

Here's how you can handle an HTTP response:

```go
// Sending a request and getting a response
resp, err := client.Do(req)
if err != nil {
    // Handle error
}
defer resp.Body.Close() // Always close the response body

// Checking the status code
if resp.StatusCode != http.StatusOK {
    // Handle non-200 status code
}

// Reading the response body
body, err := io.ReadAll(resp.Body)
if err != nil {
    // Handle error
}
```

### Convenience Functions

The `net/http` package also provides convenience functions for common HTTP operations:

```go
// Simple GET request
resp, err := http.Get("https://api.example.com/data")

// Simple POST request
resp, err := http.Post("https://api.example.com/data", "application/json", requestBody)

// HEAD, PUT, DELETE, etc. can be done using the generic method
req, err := http.NewRequest("DELETE", "https://api.example.com/data/123", nil)
resp, err := client.Do(req)
```

## Importing the Package

To use the `net/http` package in your Go code, you need to import it:

```go
import (
    "net/http"
    "io"
    "time" // If you need to set timeouts
)
```

In the next section, we'll look at how to make different types of HTTP requests using this package.
