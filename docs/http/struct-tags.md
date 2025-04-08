# Struct Tags in Go

Struct tags are a powerful feature in Go that allow you to add metadata to struct fields. They are commonly used with encoding packages like `encoding/json` to control how data is marshaled and unmarshaled.

## What Are Struct Tags?

Struct tags are string literals that follow the field declaration in a struct. They provide instructions to Go's reflection system about how to handle the field.

```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age,omitempty"`
}
```

In this example, `json:"name"` and `json:"age,omitempty"` are struct tags.

## JSON Struct Tags

The `encoding/json` package uses struct tags to determine how to map JSON data to Go struct fields. Here are the most common JSON struct tag options:

### Field Renaming

You can use struct tags to specify a different name for a field in the JSON:

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    FirstName string `json:"first_name"` // Maps to "first_name" in JSON
    LastName  string `json:"last_name"`  // Maps to "last_name" in JSON
    Age       int    `json:"age"`        // Maps to "age" in JSON
}

func main() {
    // JSON data with snake_case field names
    jsonData := []byte(`{
        "first_name": "John",
        "last_name": "Doe",
        "age": 30
    }`)

    var person Person
    err := json.Unmarshal(jsonData, &person)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("First Name:", person.FirstName)
    fmt.Println("Last Name:", person.LastName)
    fmt.Println("Age:", person.Age)

    // Marshal back to JSON
    newJSON, err := json.Marshal(person)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("JSON:", string(newJSON))
}
```

Save this code to a file named `field_renaming.go` and run it with `go run field_renaming.go`.

### Omitting Empty Fields

The `omitempty` option tells the JSON encoder to omit the field if it has an empty value (zero value or nil):

```go
package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    Name     string  `json:"name"`
    Age      int     `json:"age,omitempty"`    // Omit if zero
    Email    string  `json:"email,omitempty"`  // Omit if empty
    Address  *string `json:"address,omitempty"` // Omit if nil
}

func main() {
    // Create a user with some empty fields
    address := "123 Main St"
    user1 := User{
        Name:    "John",
        Age:     30,
        Email:   "",
        Address: &address,
    }

    user2 := User{
        Name: "Jane",
        // Age is zero, will be omitted
        // Email is not set, will be omitted
        // Address is nil, will be omitted
    }

    // Marshal to JSON
    json1, _ := json.Marshal(user1)
    json2, _ := json.Marshal(user2)

    fmt.Println("User 1:", string(json1))
    fmt.Println("User 2:", string(json2))
}
```

Save this code to a file named `omitempty.go` and run it with `go run omitempty.go`.

### Ignoring Fields

You can use the `-` tag value to tell the JSON encoder to ignore a field:

```go
package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    Name     string `json:"name"`
    Age      int    `json:"age"`
    Password string `json:"-"` // This field will be ignored during JSON marshaling
}

func main() {
    user := User{
        Name:     "John",
        Age:      30,
        Password: "secret123",
    }

    // Marshal to JSON
    jsonData, _ := json.Marshal(user)
    fmt.Println("JSON:", string(jsonData))

    // The password field will not appear in the JSON output
}
```

Save this code to a file named `ignore_field.go` and run it with `go run ignore_field.go`.

## Other Common Struct Tags

While JSON tags are the most common, Go's standard library supports other struct tags for different purposes:

### XML Tags

Similar to JSON tags, but for XML marshaling/unmarshaling:

```go
type Person struct {
    Name string `xml:"name"`
    Age  int    `xml:"age,attr"` // Will be an attribute in XML
}
```

### YAML Tags

Used by YAML libraries for marshaling/unmarshaling:

```go
type Config struct {
    ServerName string `yaml:"server_name"`
    Port       int    `yaml:"port"`
}
```

## Best Practices for Struct Tags

1. **Be consistent**: Use the same naming convention for all your JSON fields (e.g., snake_case or camelCase).
2. **Use omitempty when appropriate**: This helps keep your JSON payloads smaller.
3. **Always hide sensitive data**: Use the `-` tag for fields like passwords or API keys.
4. **Document your struct tags**: Add comments to explain non-obvious tag choices.

This concludes our section on struct tags in Go.
