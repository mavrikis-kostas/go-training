package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/murtaza-sajjad/http-client/user"
	"github.com/murtaza-sajjad/http-client/useraccount"
)

func main() {
	// ================================ Get User ID ================================
	// Get user ID from command line arguments, If no user ID is provided, return an error, e.g. go run main.go 1 here 1 is the user ID

	userID, err := userID()
	if err != nil {
		fmt.Println("Error getting user ID:", err)
		return
	}

	// ================================ Get User ================================
	// Get user information from the API, getUser function is defined below it takes a userID as an argument and returns a User struct

	userObj, err := getUser(userID)
	if err != nil {
		fmt.Println("Error getting user:", err)
		return
	}

	// ================================ Get User Accounts ================================
	// Get user accounts from the API, getUserAccounts function is defined below it takes a userID as an argument and returns a []Account struct

	accounts, err := getUserAccounts(userID)
	if err != nil {
		fmt.Println("Error getting accounts:", err)
		return
	}

	// ================================ Logs ================================
	fmt.Println("User:", userObj)
	fmt.Println("Accounts:", accounts)

	fmt.Println("================================================")

	// ================================ Display User Name ================================
	fmt.Println("User:", userObj.Attributes.Name)

	// ================================ Display Accounts ================================

	fmt.Println("Accounts:")

	totalBalance := 0
	for _, account := range accounts {
		fmt.Printf("  - %s: %d\n", account.Attributes.Name, account.Attributes.Balance)
		totalBalance += account.Attributes.Balance
	}

	// totalBalance is a variable that is the sum of all the balances of the accounts in the []Account struct
	fmt.Printf("Total Balance: %d\n", totalBalance)
}

func userID() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf("Please provide a user ID")
	}

	return os.Args[1], nil
}

func getUser(userID string) (user.User, error) {
	var resp *http.Response
	var err error
	if resp, err = http.Get("https://sample-accounts-api.herokuapp.com/users/" + userID); err != nil {
		return user.User{}, err
	}
	defer resp.Body.Close()

	var body []byte
	if body, err = io.ReadAll(resp.Body); err != nil {
		return user.User{}, err
	}

	var userObj user.User
	if err = json.Unmarshal(body, &userObj); err != nil {
		return user.User{}, err
	}

	return userObj, nil
}

func getUserAccounts(userID string) ([]useraccount.Account, error) {
	var resp *http.Response
	var err error
	if resp, err = http.Get("https://sample-accounts-api.herokuapp.com/users/" + userID + "/accounts"); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var body []byte
	if body, err = io.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	var accounts []useraccount.Account
	if err = json.Unmarshal(body, &accounts); err != nil {
		return nil, err
	}

	return accounts, nil
}
