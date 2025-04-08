# HTTP Client Exercises

This section contains practical exercises to help you apply what you've learned about making HTTP requests and parsing
JSON responses in Go.

## :material-code-braces-box: Exercise 1: Fetching and Displaying User Information

In this exercise, you'll practice making HTTP requests and parsing JSON responses by fetching user information from an
API and displaying it in a formatted way.

### Requirements

1. Make a GET request to `https://sample-accounts-api.herokuapp.com/users/1`
2. Parse the JSON response into appropriate Go structs
3. Extract and display the user's name and account IDs in the following format:
   ```
   User: <name>
   Account IDs: <id1>, <id2>, <id3>
   ```

### Expected Response Structure

The API returns a JSON response with the following structure:

```json
{
  "attributes": {
    "id": 1,
    "name": "Alice",
    "account_ids": [
      1,
      3,
      5
    ]
  }
}
```

### Hints

1. Define appropriate structs to match the nested JSON structure
2. Use the `encoding/json` package to unmarshal the response
3. Use `strings.Join` to format the account IDs for display
4. Use `strconv.Itoa` to convert integers to strings
5. Remember to handle errors properly

??? example "Click for solution"

    Here's a complete solution to the exercise:

    ```go
    package main

    import (
    	"encoding/json"
    	"fmt"
    	"io"
    	"net/http"
    	"strings"
    )

    // UserResponse represents the top-level response structure
    type UserResponse struct {
    	Attributes UserAttributes `json:"attributes"`
    }

    // UserAttributes represents the user data inside the attributes field
    type UserAttributes struct {
    	ID         int    `json:"id"`
    	Name       string `json:"name"`
    	AccountIDs []int  `json:"account_ids"`
    }

    func main() {
    	// Make the HTTP request
    	resp, err := http.Get("https://sample-accounts-api.herokuapp.com/users/1")
    	if err != nil {
    		fmt.Println("Error making HTTP request:", err)
    		return
    	}
    	defer resp.Body.Close()

    	// Check if the request was successful
    	if resp.StatusCode != http.StatusOK {
    		fmt.Printf("Unexpected status code: %v\n", resp.StatusCode)
    		return
    	}

    	// Read the response body
    	body, err := io.ReadAll(resp.Body)
    	if err != nil {
    		fmt.Println("Error reading response body:", err)
    		return
    	}

    	// Parse the JSON response
    	var userResp UserResponse
    	err = json.Unmarshal(body, &userResp)
    	if err != nil {
    		fmt.Println("Error parsing JSON:", err)
    		return
    	}

    	// Convert account IDs to strings for display
    	accountIDStrings := make([]string, len(userResp.Attributes.AccountIDs))
    	for i, id := range userResp.Attributes.AccountIDs {
    		accountIDStrings[i] = strconv.Itoa(id)
    	}
    	accountIDs := strings.Join(accountIDStrings, ", ")

    	// Display the user information
    	fmt.Printf("User: %v\n", userResp.Attributes.Name)
    	fmt.Printf("Account IDs: %v\n", accountIDs)
    }
    ```

## :material-code-braces-box: Exercise 2: Displaying User Account Information

In this exercise, you'll create a program that takes a user ID as input and returns the user's name, account list, and
balances.

### Requirements

1. Take a user ID as input from standard input
2. Make a GET request to `https://sample-accounts-api.herokuapp.com/users/{id}` to get the user's name
3. Make a GET request to `https://sample-accounts-api.herokuapp.com/users/{id}/accounts` to get the user's accounts and
   balances
4. Display the user's name, account names, balances, and total balance in the following format:
   ```
   User: <name>
   Accounts:
     - <account_name>: <balance>
     - <account_name>: <balance>
     ...
   Total Balance: <total_balance>
   ```

### Expected Response Structures

The `users/{id}` API endpoint returns a JSON object with the following structure:

```json
{
  "attributes": {
    "id": 1,
    "name": "Alice",
    "account_ids": [
      1,
      3,
      5
    ]
  }
}
```

The `users/{id}/accounts` API endpoint returns a JSON array of account objects with the following structure:

```json
[
  {
    "attributes": {
      "id": 1,
      "user_id": 1,
      "name": "A銀行",
      "balance": 20000
    }
  },
  {
    "attributes": {
      "id": 3,
      "user_id": 1,
      "name": "C信用金庫",
      "balance": 120000
    }
  },
  {
    "attributes": {
      "id": 5,
      "user_id": 1,
      "name": "E銀行",
      "balance": 5000
    }
  }
]
```

### Hints

1. Use `fmt.Scanln` to read the user ID from standard input
2. Use `strconv.Atoi` to validate that the user ID is a number
3. Try to split the code into functions for better organization
4. Functions can return custom errors by using the `fmt.Errorf` function

??? example "Click for solution"

    Here's a complete solution to the exercise:

    ```go
    package main
    
    import (
    	"encoding/json"
    	"fmt"
    	"io"
    	"net/http"
    	"strconv"
    )
    
    // UserResponse represents the top-level response structure for user info
    type UserResponse struct {
    	Attributes UserAttributes `json:"attributes"`
    }
    
    // UserAttributes represents the user data inside the attributes field
    type UserAttributes struct {
    	ID         int    `json:"id"`
    	Name       string `json:"name"`
    	AccountIDs []int  `json:"account_ids"`
    }
    
    // AccountResponse represents a single account in the accounts array
    type AccountResponse struct {
    	Attributes AccountAttributes `json:"attributes"`
    }
    
    // AccountAttributes represents the account data inside the attributes field
    type AccountAttributes struct {
    	ID      int    `json:"id"`
    	UserID  int    `json:"user_id"`
    	Name    string `json:"name"`
    	Balance int    `json:"balance"`
    }
    
    func main() {
    	var userID string
    
    	// Prompt for user ID and validate input
    	for {
    		fmt.Print("Enter a user ID: ")
    		fmt.Scanln(&userID)
    
    		if _, err := strconv.Atoi(userID); err == nil {
    			break
    		}
    		fmt.Println("Invalid input. Please enter a valid user ID.")
    	}
    
    	// Get user information
    	userName, err := getUserName(userID)
    	if err != nil {
    		fmt.Println("Error getting user information:", err)
    		return
    	}
    
    	// Get user accounts
    	accounts, err := getUserAccounts(userID)
    	if err != nil {
    		fmt.Println("Error getting user accounts:", err)
    		return
    	}
    
    	// Calculate total balance
    	totalBalance := 0
    	for _, account := range accounts {
    		totalBalance += account.Attributes.Balance
    	}
    
    	// Display user information
    	fmt.Printf("User: %v\n", userName)
    	fmt.Println("Accounts:")
    	for _, account := range accounts {
    		fmt.Printf("  - %v: %v\n", account.Attributes.Name, account.Attributes.Balance)
    	}
    	fmt.Printf("Total Balance: %v\n", totalBalance)
    }
    
    // getUserName fetches the user's name from the API
    func getUserName(userID string) (string, error) {
    	// Make the HTTP request
    	resp, err := http.Get("https://sample-accounts-api.herokuapp.com/users/" + userID)
    	if err != nil {
    		return "", fmt.Errorf("error making HTTP request: %w", err)
    	}
    	defer resp.Body.Close()
    
    	// Check if the request was successful
    	if resp.StatusCode != http.StatusOK {
    		return "", fmt.Errorf("unexpected status code: %v", resp.StatusCode)
    	}
    
    	// Read the response body
    	body, err := io.ReadAll(resp.Body)
    	if err != nil {
    		return "", fmt.Errorf("error reading response body: %w", err)
    	}
    
    	// Parse the JSON response
    	var userResp UserResponse
    	err = json.Unmarshal(body, &userResp)
    	if err != nil {
    		return "", fmt.Errorf("error parsing JSON: %w", err)
    	}
    
    	return userResp.Attributes.Name, nil
    }
    
    // getUserAccounts fetches the user's accounts from the API
    func getUserAccounts(userID string) ([]AccountResponse, error) {
    	// Make the HTTP request
    	resp, err := http.Get("https://sample-accounts-api.herokuapp.com/users/" + userID + "/accounts")
    	if err != nil {
    		return nil, fmt.Errorf("error making HTTP request: %w", err)
    	}
    	defer resp.Body.Close()
    
    	// Check if the request was successful
    	if resp.StatusCode != http.StatusOK {
    		return nil, fmt.Errorf("unexpected status code: %v", resp.StatusCode)
    	}
    
    	// Read the response body
    	body, err := io.ReadAll(resp.Body)
    	if err != nil {
    		return nil, fmt.Errorf("error reading response body: %w", err)
    	}
    
    	// Parse the JSON response
    	var accounts []AccountResponse
    	err = json.Unmarshal(body, &accounts)
    	if err != nil {
    		return nil, fmt.Errorf("error parsing JSON: %w", err)
    	}
    
    	return accounts, nil
    }
    ```
