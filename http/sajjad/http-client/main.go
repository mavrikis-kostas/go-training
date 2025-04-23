package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/murtaza-sajjad/http-client/models"
)

func main() {
	// Check if user ID is provided as command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <user_id>")
		return
	}

	userID := os.Args[1]

	// Validate that the user ID is a number
	_, err := strconv.Atoi(userID)
	if err != nil {
		fmt.Println("Invalid user ID. Please provide a numeric user ID:", err)
		return
	}

	// Construct URL with the dynamic user ID
	userURL := fmt.Sprintf("https://sample-accounts-api.herokuapp.com/users/%s", userID)
	resp, err := http.Get(userURL)
	if err != nil {
		fmt.Printf("error fetching user %s: %v\n", userID, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("unexpected status code fetching user %s: %d\n", userID, resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading user response body: %v\n", err)
		return
	}

	// Unmarshal the JSON response into a User struct
	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		fmt.Printf("error parsing user data: %v\n", err)
		return
	}

	// Display user name
	fmt.Println("User:", user.Attributes.Name)

	// Fetch and display account details
	accountsURL := fmt.Sprintf("https://sample-accounts-api.herokuapp.com/users/%s/accounts", userID)
	accountsResp, err := http.Get(accountsURL)
	if err != nil {
		fmt.Printf("error fetching accounts for user %s: %v\n", userID, err)
		return
	}
	defer accountsResp.Body.Close()

	accountsBody, err := io.ReadAll(accountsResp.Body)
	if err != nil {
		fmt.Printf("error reading accounts response body: %v\n", err)
		return
	}

	var accounts []models.Account
	if err = json.Unmarshal(accountsBody, &accounts); err != nil {
		fmt.Printf("error parsing accounts data: %v\n", err)
		return
	}

	// Display accounts and calculate total balance
	fmt.Println("Accounts:")
	totalBalance := 0
	for _, account := range accounts {
		fmt.Printf("  - %s: %d\n", account.Attributes.Name, account.Attributes.Balance)
		totalBalance += account.Attributes.Balance
	}

	// Display total balance
	fmt.Printf("Total Balance: %d\n", totalBalance)
}
