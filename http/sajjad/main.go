package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// {"attributes":{"id":1,"name":"Alice","account_ids":[1,3,5]}}
type User struct {

	// Attributes struct {
	// 	Id         int    `json:"id"`
	// 	Name       string `json:"name"`
	// 	AccountIds []int  `json:"account_ids"`
	// } `json:"attributes"`

	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	AccountIds []int  `json:"account_ids"`
}

func main() {
	resp, err := http.Get("https://sample-accounts-api.herokuapp.com/users/1")

	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected status code:", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Unmarshal the JSON response into a User struct
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	accountIDs := make([]string, len(user.Attributes.AccountIds))
	for i, accountID := range user.Attributes.AccountIds {
		accountIDs[i] = strconv.Itoa(accountID)
	}
	fmt.Println("User:", user.Attributes.Name)

	fmt.Println("Account IDs:", strings.Join(accountIDs, ","))
}
