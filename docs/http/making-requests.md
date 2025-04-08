# Making HTTP Requests

In this section, we'll explore how to make different types of HTTP requests using Go's `net/http` package.

## Making a GET Request

GET requests are used to retrieve data from a server. Here's how to make a simple GET request:

```go
package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {
    // Create a new GET request
    resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
    if err != nil {
        fmt.Println("Error making request:", err)
        return
    }
    defer resp.Body.Close()

    // Check if the request was successful
    if resp.StatusCode != http.StatusOK {
        fmt.Println("Unexpected status code:", resp.StatusCode)
        return
    }

    // Read and print the response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    fmt.Println("Response body:")
    fmt.Println(string(body))
}
```

Save this code to a file named `get_request.go` and run it with `go run get_request.go`.

## Making a POST Request

POST requests are used to send data to a server to create or update a resource. Here's how to make a POST request with a JSON payload:

```go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

func main() {
    // Create the data to send
    data := map[string]interface{}{
        "title":  "foo",
        "body":   "bar",
        "userId": 1,
    }

    // Convert data to JSON
    jsonData, err := json.Marshal(data)
    if err != nil {
        fmt.Println("Error marshaling JSON:", err)
        return
    }

    // Create a new POST request with the JSON data
    resp, err := http.Post(
        "https://jsonplaceholder.typicode.com/posts",
        "application/json",
        bytes.NewBuffer(jsonData),
    )
    if err != nil {
        fmt.Println("Error making request:", err)
        return
    }
    defer resp.Body.Close()

    // Check if the request was successful
    if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
        fmt.Println("Unexpected status code:", resp.StatusCode)
        return
    }

    // Read and print the response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    fmt.Println("Response body:")
    fmt.Println(string(body))
}
```

Save this code to a file named `post_request.go` and run it with `go run post_request.go`.

## Making a Request with Custom Headers

Sometimes you need to add custom headers to your HTTP requests, such as authentication tokens or content type specifications:

```go
package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {
    // Create a new request
    req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    // Add custom headers
    req.Header.Add("Accept", "application/json")
    req.Header.Add("User-Agent", "Go-HTTP-Client/1.0")
    // For authentication, you might add something like:
    // req.Header.Add("Authorization", "Bearer your-token-here")

    // Create a client and send the request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error making request:", err)
        return
    }
    defer resp.Body.Close()

    // Check if the request was successful
    if resp.StatusCode != http.StatusOK {
        fmt.Println("Unexpected status code:", resp.StatusCode)
        return
    }

    // Read and print the response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    fmt.Println("Response body:")
    fmt.Println(string(body))
}
```

Save this code to a file named `custom_headers.go` and run it with `go run custom_headers.go`.

In the next section, we'll learn how to parse JSON responses from APIs.
