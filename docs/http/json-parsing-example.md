# JSON Parsing Example

This example demonstrates simple techniques for parsing JSON data in Go using struct tags.

## The Code

Save this code to a file named `json_parsing.go`:

```go
package main

import (
	"encoding/json"
	"fmt"
)

// Product represents a simple product with basic fields
type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	InStock   bool    `json:"in_stock"`
	CreatedAt string  `json:"created_at"`
}

// User represents a user with a nested address
type User struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

// Address represents a simple address
type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	Zip    string `json:"zip"`
}

func main() {
	// Example 1: Basic JSON parsing with struct tags
	fmt.Println("Example 1: Basic JSON parsing")
	productJSON := `{
        "id": 101,
        "name": "Laptop",
        "price": 999.99,
        "in_stock": true,
        "created_at": "2023-01-15"
    }`

	var product Product
	json.Unmarshal([]byte(productJSON), &product)

	fmt.Printf("Product: %s (ID: %d)\n", product.Name, product.ID)
	fmt.Printf("Price: $%.2f, In Stock: %t\n", product.Price, product.InStock)
	fmt.Printf("Created: %s\n\n", product.CreatedAt)

	// Example 2: Parsing JSON arrays
	fmt.Println("Example 2: Parsing JSON arrays")
	productsJSON := `[
        {"id": 101, "name": "Laptop", "price": 999.99, "in_stock": true},
        {"id": 102, "name": "Phone", "price": 699.99, "in_stock": true},
        {"id": 103, "name": "Headphones", "price": 199.99, "in_stock": false}
    ]`

	var products []Product
	json.Unmarshal([]byte(productsJSON), &products)

	fmt.Printf("Found %d products:\n", len(products))
	for _, p := range products {
		fmt.Printf("- %s: $%.2f\n", p.Name, p.Price)
	}
	fmt.Println()

	// Example 3: Parsing nested JSON
	fmt.Println("Example 3: Parsing nested JSON")
	userJSON := `{
        "name": "John Doe",
        "email": "john@example.com",
        "address": {
            "street": "123 Main St",
            "city": "Anytown",
            "zip": "12345"
        }
    }`

	var user User
	json.Unmarshal([]byte(userJSON), &user)

	fmt.Printf("User: %s (%s)\n", user.Name, user.Email)
	fmt.Printf("Address: %s, %s %s\n\n", user.Address.Street, user.Address.City, user.Address.Zip)
}
```

## Running the Example

To run this example:

1. Save the code to a file named `json_parsing.go`
2. Open a terminal and navigate to the directory containing the file
3. Run the command: `go run json_parsing.go`

## What This Example Demonstrates

This example demonstrates:

1. Basic JSON parsing using struct tags
2. Parsing JSON arrays into slices of structs
3. Handling nested JSON objects with nested structs

These techniques cover the most common JSON parsing needs in Go applications.
