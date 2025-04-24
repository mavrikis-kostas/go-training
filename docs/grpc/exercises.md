# gRPC Exercises

This section contains practical exercises to help you apply what you've learned about gRPC in Go.

## :material-code-braces-box: Exercise 1: Basic gRPC Service

In this exercise, you'll practice creating a simple gRPC service that provides information about books.

### Requirements

1. Create a `.proto` file that defines a `BookService` with the following methods:
    - `GetBook`: Takes a book ID and returns book details (unary RPC)
    - `ListBooks`: Returns details of all books (server streaming RPC)

2. The `Book` message should include:
    - `id` (int32)
    - `title` (string)
    - `author` (string)
    - `year` (int32)

3. Implement a gRPC server that:
    - Stores a collection of books in memory
    - Implements the `GetBook` and `ListBooks` methods

4. Implement a gRPC client that:
    - Calls `GetBook` with a specific ID
    - Calls `ListBooks` and prints all books

### Hints

1. Start by defining your `.proto` file
2. Generate the Go code using `protoc`
3. Implement the server with a simple in-memory storage
4. Implement the client to call both methods

??? example "Click for solution"

    **book.proto**:
    ```protobuf
    syntax = "proto3";

    package bookservice;

    option go_package = "example.com/bookservice";

    // Book service definition
    service BookService {
      // Get a book by ID
      rpc GetBook (GetBookRequest) returns (Book);
      
      // List all books
      rpc ListBooks (ListBooksRequest) returns (stream Book);
    }

    // Request message for GetBook
    message GetBookRequest {
      int32 id = 1;
    }

    // Request message for ListBooks (empty as we want all books)
    message ListBooksRequest {}

    // Book message
    message Book {
      int32 id = 1;
      string title = 2;
      string author = 3;
      int32 year = 4;
    }
    ```

    **server/main.go**:
    ```go
    package main

    import (
        "context"
        "fmt"
        "log"
        "net"

        "google.golang.org/grpc"
        "google.golang.org/grpc/codes"
        "google.golang.org/grpc/status"

        pb "example.com/bookservice/bookservice"
    )

    // server is used to implement the BookService
    type server struct {
        pb.UnimplementedBookServiceServer
        books []*pb.Book
    }

    // GetBook implements the GetBook RPC method
    func (s *server) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.Book, error) {
        for _, book := range s.books {
            if book.Id == req.Id {
                return book, nil
            }
        }
        return nil, status.Errorf(codes.NotFound, "book with ID %d not found", req.Id)
    }

    // ListBooks implements the ListBooks RPC method
    func (s *server) ListBooks(req *pb.ListBooksRequest, stream pb.BookService_ListBooksServer) error {
        for _, book := range s.books {
            if err := stream.Send(book); err != nil {
                return err
            }
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

        // Create sample books
        books := []*pb.Book{
            {Id: 1, Title: "The Go Programming Language", Author: "Alan A. A. Donovan and Brian W. Kernighan", Year: 2015},
            {Id: 2, Title: "Go in Action", Author: "William Kennedy", Year: 2015},
            {Id: 3, Title: "Concurrency in Go", Author: "Katherine Cox-Buday", Year: 2017},
        }

        // Register our service with the gRPC server
        pb.RegisterBookServiceServer(s, &server{books: books})

        // Start serving requests
        fmt.Println("Server started on :50051")
        if err := s.Serve(lis); err != nil {
            log.Fatalf("Failed to serve: %v", err)
        }
    }
    ```

    **client/main.go**:
    ```go
    package main

    import (
        "context"
        "io"
        "log"
        "time"

        "google.golang.org/grpc"
        "google.golang.org/grpc/credentials/insecure"

        pb "example.com/bookservice/bookservice"
    )

    func main() {
        // Set up a connection to the server
        conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
        if err != nil {
            log.Fatalf("Failed to connect: %v", err)
        }
        defer conn.Close()

        // Create a client
        client := pb.NewBookServiceClient(conn)

        // Set a timeout for our API calls
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()

        // Call GetBook with ID 2
        book, err := client.GetBook(ctx, &pb.GetBookRequest{Id: 2})
        if err != nil {
            log.Fatalf("GetBook failed: %v", err)
        }
        log.Printf("Book: ID=%d, Title=%s, Author=%s, Year=%d", 
            book.Id, book.Title, book.Author, book.Year)

        // Call ListBooks
        stream, err := client.ListBooks(ctx, &pb.ListBooksRequest{})
        if err != nil {
            log.Fatalf("ListBooks failed: %v", err)
        }

        log.Println("All books:")
        for {
            book, err := stream.Recv()
            if err == io.EOF {
                break
            }
            if err != nil {
                log.Fatalf("Error while receiving: %v", err)
            }
            log.Printf("  ID=%d, Title=%s, Author=%s, Year=%d", 
                book.Id, book.Title, book.Author, book.Year)
        }
    }
    ```

## :material-code-braces-box: Exercise 2: Client Streaming RPC

In this exercise, you'll practice implementing a client streaming RPC method.

### Requirements

1. Extend the `BookService` from Exercise 1 to add a new method:
    - `AddBooks`: Takes a stream of books from the client and returns the total number of books added (client streaming
      RPC)

2. Update your server to:
    - Implement the `AddBooks` method
    - Add the received books to the in-memory collection

3. Update your client to:
    - Call `AddBooks` with several new books
    - Print the total number of books added

### Hints

1. Update your `.proto` file to add the new method
2. Regenerate the Go code using `protoc`
3. Implement the new method in your server
4. Update your client to call the new method

??? example "Click for solution"

    **Updated book.proto** (add this method to the BookService):
    ```protobuf
    // Add multiple books (client streaming)
    rpc AddBooks (stream Book) returns (AddBooksResponse);
    
    // Response message for AddBooks
    message AddBooksResponse {
      int32 total_added = 1;
    }
    ```

    **Updated server/main.go** (add this method to the server struct):
    ```go
    // AddBooks implements the AddBooks RPC method
    func (s *server) AddBooks(stream pb.BookService_AddBooksServer) error {
        var addedCount int32 = 0
        
        for {
            book, err := stream.Recv()
            if err == io.EOF {
                // End of stream, return the response
                return stream.SendAndClose(&pb.AddBooksResponse{
                    TotalAdded: addedCount,
                })
            }
            if err != nil {
                return err
            }
            
            // Generate a new ID (simple approach)
            maxID := int32(0)
            for _, b := range s.books {
                if b.Id > maxID {
                    maxID = b.Id
                }
            }
            book.Id = maxID + 1
            
            // Add the book to our collection
            s.books = append(s.books, book)
            addedCount++
            
            log.Printf("Added book: %s by %s", book.Title, book.Author)
        }
    }
    ```

    **Updated client/main.go** (add this after the ListBooks call):
    ```go
    // Call AddBooks with several new books
    addBooksStream, err := client.AddBooks(ctx)
    if err != nil {
        log.Fatalf("AddBooks stream creation failed: %v", err)
    }
    
    // Send several books
    newBooks := []*pb.Book{
        {Title: "Effective Go", Author: "Google", Year: 2009},
        {Title: "Go Programming Blueprints", Author: "Mat Ryer", Year: 2016},
        {Title: "Go Web Programming", Author: "Sau Sheong Chang", Year: 2016},
    }
    
    for _, book := range newBooks {
        if err := addBooksStream.Send(book); err != nil {
            log.Fatalf("Failed to send book: %v", err)
        }
    }
    
    // Close the stream and get the response
    resp, err := addBooksStream.CloseAndRecv()
    if err != nil {
        log.Fatalf("Failed to receive response: %v", err)
    }
    
    log.Printf("Added %d books successfully", resp.TotalAdded)
    ```

## :material-code-braces-box: Exercise 3: Bidirectional Streaming RPC

In this exercise, you'll practice implementing a bidirectional streaming RPC method.

### Requirements

1. Extend the `BookService` from Exercise 2 to add a new method:
    - `SearchBooks`: Takes a stream of search queries from the client and returns a stream of matching books (
      bidirectional streaming RPC)

2. Update your server to:
    - Implement the `SearchBooks` method
    - For each search query, return all books that contain the query string in their title or author

3. Update your client to:
    - Call `SearchBooks` with several search queries
    - Print all matching books for each query

### Hints

1. Update your `.proto` file to add the new method
2. Regenerate the Go code using `protoc`
3. Implement the new method in your server
4. Update your client to call the new method

??? example "Click for solution"

    **Updated book.proto** (add this to the BookService):
    ```protobuf
    // Search for books (bidirectional streaming)
    rpc SearchBooks (stream SearchRequest) returns (stream SearchResponse);
    
    // Request message for SearchBooks
    message SearchRequest {
      string query = 1;
    }
    
    // Response message for SearchBooks
    message SearchResponse {
      string query = 1;
      Book book = 2;
    }
    ```

    **Updated server/main.go** (add this method to the server struct):
    ```go
    // SearchBooks implements the SearchBooks RPC method
    func (s *server) SearchBooks(stream pb.BookService_SearchBooksServer) error {
        for {
            req, err := stream.Recv()
            if err == io.EOF {
                return nil
            }
            if err != nil {
                return err
            }
            
            query := strings.ToLower(req.Query)
            log.Printf("Received search query: %s", query)
            
            // Search for matching books
            for _, book := range s.books {
                if strings.Contains(strings.ToLower(book.Title), query) || 
                   strings.Contains(strings.ToLower(book.Author), query) {
                    // Send matching book to the client
                    response := &pb.SearchResponse{
                        Query: req.Query,
                        Book: book,
                    }
                    if err := stream.Send(response); err != nil {
                        return err
                    }
                }
            }
        }
    }
    ```

    **Updated client/main.go** (add this after the AddBooks call):
    ```go
    // Call SearchBooks with several queries
    searchStream, err := client.SearchBooks(ctx)
    if err != nil {
        log.Fatalf("SearchBooks stream creation failed: %v", err)
    }
    
    // Create a wait group to wait for the goroutines to finish
    var wg sync.WaitGroup
    wg.Add(2)
    
    // Start a goroutine to send search queries
    go func() {
        defer wg.Done()
        
        queries := []string{"Go", "Programming", "2016"}
        for _, query := range queries {
            log.Printf("Searching for: %s", query)
            if err := searchStream.Send(&pb.SearchRequest{Query: query}); err != nil {
                log.Fatalf("Failed to send query: %v", err)
                return
            }
            time.Sleep(500 * time.Millisecond) // Small delay between queries
        }
        
        // Close the send direction of the stream
        searchStream.CloseSend()
    }()
    
    // Start a goroutine to receive search results
    go func() {
        defer wg.Done()
        
        for {
            resp, err := searchStream.Recv()
            if err == io.EOF {
                break
            }
            if err != nil {
                log.Fatalf("Failed to receive response: %v", err)
                return
            }
            
            log.Printf("Result for query '%s': %s by %s (%d)", 
                resp.Query, resp.Book.Title, resp.Book.Author, resp.Book.Year)
        }
    }()
    
    // Wait for both goroutines to complete
    wg.Wait()
    ```

    Don't forget to add these imports if they're not already present:
    ```go
    import (
        "strings"
        "sync"
    )
    ```