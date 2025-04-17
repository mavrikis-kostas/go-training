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

type AccountsResponse struct {
	AccountsAttributes AccountAttributes `json:"attributes"`
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
	accounts, err := getAccounts(userId)

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println("user: ", user)
	printAccounts(accounts)
}

func getUser(id int) (string, error) {
	resp, err := http.Get(URL + strconv.Itoa(id))

	if err != nil {
		fmt.Println("Error making request: ", err)
		return "", fmt.Errorf("error in making request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("response error: ", resp.StatusCode)
		return "", fmt.Errorf("response error: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	var res UserResponse
	err = json.Unmarshal(body, &res)

	if err != nil {
		fmt.Println("error in parsing json: ", err)
		return "", fmt.Errorf("error in parsing json: %v", err)
	}

	return res.UserAttributes.Name, nil
}

func getAccounts(id int) ([]AccountAttributes, error) {
	resp, err := http.Get(URL + strconv.Itoa(id) + "/accounts")

	if err != nil {
		fmt.Println("Error making request: ", err)
		return nil, fmt.Errorf("error in making request: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		return nil, fmt.Errorf("error in reading response body: %v", err)
	}

	var res []AccountsResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("error in parsing json: ", err)
		return nil, fmt.Errorf("error in parsing json: %v", err)
	}

	var accounts []AccountAttributes = []AccountAttributes{}

	for _, account := range res {
		accounts = append(accounts, account.AccountsAttributes)
	}

	return accounts, nil
}

func printAccounts(accounts []AccountAttributes) {
	var totalBalance int = 0

	for _, account := range accounts {
		fmt.Println(account.Name, ": ", account.Balance)
		totalBalance += account.Balance
	}
	fmt.Println("Total balance: ", totalBalance)
}
