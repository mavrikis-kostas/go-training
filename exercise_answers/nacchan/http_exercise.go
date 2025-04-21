package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type UserAttributes struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	AccountIds []int  `json:"account_ids"`
}

type AccountAttributes struct {
	Id      int    `json:"id"`
	UserId  int    `json:"user_id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

type UserResponse struct {
	UserAttributes UserAttributes `json:"attributes"`
}

type AccountResponse struct {
	AccountAttributes AccountAttributes `json:"attributes"`
}

const URL = "https://sample-accounts-api.herokuapp.com/users/"

func main() {
	var userId int
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter user id: ")
		userIdString, err := reader.ReadString('\n')
		userIdString = strings.TrimSpace(userIdString)
		userId, err = strconv.Atoi(userIdString)
		if err != nil {
			fmt.Println("please enter a valid integer")
			continue
		}
		break
	}

	user, err := getUser(userId)
	if err != nil {
		fmt.Println("Error getting user: ", err)
		return
	}

	accounts, err := getAccounts(userId)
	if err != nil {
		fmt.Println("Error getting accounts: ", err)
		return
	}

	fmt.Println("User: ", user)
	fmt.Println("Accounts: ")
	printAccounts(accounts)
}

func getUser(id int) (string, error) {
	resp, err := http.Get(URL + strconv.Itoa(id))

	if err != nil {
		return "", fmt.Errorf("error in making request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("response error: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error in reading response body: %v", err)
	}

	var res UserResponse
	err = json.Unmarshal(body, &res)

	if err != nil {
		return "", fmt.Errorf("error in parsing json: %v", err)
	}

	return res.UserAttributes.Name, nil
}

func getAccounts(id int) ([]AccountAttributes, error) {
	resp, err := http.Get(URL + strconv.Itoa(id) + "/accounts")

	if err != nil {
		return nil, fmt.Errorf("error in making request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response error: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error in reading response body: %v", err)
	}

	var res []AccountResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, fmt.Errorf("error in parsing json: %v", err)
	}

	var accounts []AccountAttributes = []AccountAttributes{}

	for _, account := range res {
		accounts = append(accounts, account.AccountAttributes)
	}

	return accounts, nil
}

func printAccounts(accounts []AccountAttributes) {
	var totalBalance int = 0

	for _, account := range accounts {
		fmt.Println(" -", account.Name, ": ", account.Balance)
		totalBalance += account.Balance
	}
	fmt.Println("Total balance: ", totalBalance)
}
