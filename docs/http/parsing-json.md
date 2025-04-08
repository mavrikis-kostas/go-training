# Parsing JSON Responses

Most modern APIs return data in JSON format. In this section, we'll learn how to parse JSON responses in Go using the `encoding/json` package and struct tags.

## Basic JSON Parsing with Maps

Here's a simple example of parsing a JSON string into a map:

```go
package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    // Sample JSON data
    jsonData := `{"id": 1, "title": "Hello", "completed": false}`

    // Parse JSON into a map
    var result map[string]interface{}
    json.Unmarshal([]byte(jsonData), &result)

    // Access the data
    fmt.Println("ID:", result["id"])
    fmt.Println("Title:", result["title"])
    fmt.Println("Completed:", result["completed"])
}
```

## Parsing JSON with Struct Tags

Using struct tags makes JSON parsing more robust and type-safe:

```go
package main

import (
    "encoding/json"
    "fmt"
)

// Define a struct with JSON struct tags
type Todo struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}

func main() {
    // Sample JSON data
    jsonData := `{"id": 1, "title": "Hello", "completed": false}`

    // Parse JSON into a struct
    var todo Todo
    json.Unmarshal([]byte(jsonData), &todo)

    // Access the data using struct fields
    fmt.Println("ID:", todo.ID)
    fmt.Println("Title:", todo.Title)
    fmt.Println("Completed:", todo.Completed)
}
```

## Parsing JSON Arrays

Here's how to parse a JSON array into a slice of structs:

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Todo struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}

func main() {
    // Sample JSON array
    jsonData := `[
        {"id": 1, "title": "Task 1", "completed": false},
        {"id": 2, "title": "Task 2", "completed": true}
    ]`

    // Parse JSON array into a slice of structs
    var todos []Todo
    json.Unmarshal([]byte(jsonData), &todos)

    // Loop through the todos
    for i, todo := range todos {
        fmt.Printf("Todo %d: %s (Completed: %t)\n", 
            i+1, todo.Title, todo.Completed)
    }
}
```

## Handling Nested JSON

For nested JSON structures, you can use nested structs with struct tags:

```go
package main

import (
    "encoding/json"
    "fmt"
)

// Define structs for nested JSON
type User struct {
    Name    string  `json:"name"`
    Email   string  `json:"email"`
    Address Address `json:"address"`
}

type Address struct {
    Street string `json:"street"`
    City   string `json:"city"`
    Zip    string `json:"zip"`
}

func main() {
    // Sample nested JSON
    jsonData := `{
        "name": "John Doe",
        "email": "john@example.com",
        "address": {
            "street": "123 Main St",
            "city": "Anytown",
            "zip": "12345"
        }
    }`

    // Parse nested JSON
    var user User
    json.Unmarshal([]byte(jsonData), &user)

    // Access the nested data
    fmt.Println("Name:", user.Name)
    fmt.Println("Email:", user.Email)
    fmt.Println("Street:", user.Address.Street)
    fmt.Println("City:", user.Address.City)
    fmt.Println("Zip:", user.Address.Zip)
}
```

These examples demonstrate how to use struct tags to parse JSON data in Go. Struct tags make it easy to map between JSON field names and Go struct field names, even when they don't match exactly.
