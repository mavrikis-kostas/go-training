package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UserResponse struct {
	Attributes User `json:"attributes"`
}

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	AccountIDs []int  `json:"account_ids"`
}

type AccountResponse struct {
	Attributes Account `json:"attributes"`
}

type Account struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

func getUser(userID int) (User, error) {
	url := fmt.Sprintf("https://sample-accounts-api.herokuapp.com/users/%d", userID)
	resp, err := http.Get(url)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return User{}, err
	}

	var userResponse UserResponse
	err = json.Unmarshal(body, &userResponse)
	if err != nil {
		return User{}, err
	}

	return userResponse.Attributes, nil
}

func getAccounts(userID int) ([]Account, error) {
	url := fmt.Sprintf("https://sample-accounts-api.herokuapp.com/users/%d/accounts", userID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var accountResponses []AccountResponse
	err = json.Unmarshal(body, &accountResponses)
	if err != nil {
		return nil, err
	}

	accounts := make([]Account, len(accountResponses))
	for i, accountResponse := range accountResponses {
		accounts[i] = accountResponse.Attributes
	}

	return accounts, nil
}

func displayUser(user User, accounts []Account) {
	fmt.Printf("User: %s\n", user.Name)
	fmt.Println("Accounts:")
	totalBalance := 0
	for _, account := range accounts {
		fmt.Printf("  - %s: %d\n", account.Name, account.Balance)
		totalBalance += account.Balance
	}
	fmt.Printf("Total Balance: %d\n", totalBalance)
}

func main() {
	var userID int
	_, err := fmt.Scanln(&userID)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	user, err := getUser(userID)
	if err != nil {
		fmt.Println("error found")
		return
	}

	accounts, err := getAccounts(userID)
	if err != nil {
		fmt.Println("error found")
		return
	}

	displayUser(user, accounts)
}
