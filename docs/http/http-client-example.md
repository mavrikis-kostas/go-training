# Simple HTTP Client Example

This example demonstrates a complete HTTP client application that fetches data from an API, processes it, and displays the results.

## The Code

Save this code to a file named `http_client.go`:

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Post represents a blog post from the JSONPlaceholder API
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	fmt.Println("HTTP Client Example")
	fmt.Println("==================")

	// Fetch all posts
	fmt.Println("\nFetching all posts...")
	posts, err := fetchPosts()
	if err != nil {
		fmt.Printf("Error fetching posts: %v\n", err)
		return
	}
	fmt.Printf("Fetched %d posts\n", len(posts))

	// Print the first 3 posts
	fmt.Println("\nFirst 3 posts:")
	for i, post := range posts {
		if i >= 3 {
			break
		}
		fmt.Printf("Post %d: %s\n", post.ID, post.Title)
	}

	// Fetch a single post
	fmt.Println("\nFetching post with ID 1...")
	post, err := fetchPost(1)
	if err != nil {
		fmt.Printf("Error fetching post: %v\n", err)
		return
	}
	fmt.Printf("Post title: %s\n", post.Title)
	fmt.Printf("Post body: %s\n", post.Body)

	// Create a new post
	fmt.Println("\nCreating a new post...")
	newPost := Post{
		UserID: 1,
		Title:  "My New Post",
		Body:   "This is the content of my new post.",
	}
	createdPost, err := createPost(newPost)
	if err != nil {
		fmt.Printf("Error creating post: %v\n", err)
		return
	}
	fmt.Printf("Created post with ID %d\n", createdPost.ID)
	fmt.Printf("Post title: %s\n", createdPost.Title)
	fmt.Printf("Post body: %s\n", createdPost.Body)
}

// fetchPosts retrieves posts from the JSONPlaceholder API
func fetchPosts() ([]Post, error) {
	// Create a client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Make the request
	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Parse the JSON
	var posts []Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return posts, nil
}

// fetchPost retrieves a single post by ID
func fetchPost(id int) (*Post, error) {
	// Create a client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Make the request
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id)
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Parse the JSON
	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return &post, nil
}

// createPost creates a new post
func createPost(post Post) (*Post, error) {
	// Create a client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Convert the post to JSON
	postJSON, err := json.Marshal(post)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %w", err)
	}

	// Create a request with the JSON body
	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Parse the JSON
	var createdPost Post
	err = json.Unmarshal(body, &createdPost)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return &createdPost, nil
}

```

## Running the Example

To run this example:

1. Save the code to a file named `http_client.go`
2. Open a terminal and navigate to the directory containing the file
3. Run the command: `go run http_client.go`

You should see output showing the results of fetching posts, fetching a single post, and creating a new post.

## What This Example Demonstrates

This example demonstrates:

1. Making GET and POST requests to an API
2. Setting timeouts to prevent hanging requests
3. Handling HTTP status codes
4. Parsing JSON responses into Go structs
5. Using struct tags for JSON mapping
6. Proper error handling
7. Creating JSON from Go structs

These are all essential skills for building applications that interact with web APIs.