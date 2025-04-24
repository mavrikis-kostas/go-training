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

// here its used by json := exported
type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	AccountIDs []int  `json:"account_ids"`
}
// array_of_ints = {array_of_strings}.map(:to_s)
// convert an array into something else
// 1. create a second array variable
// arrayOfInts := make([]int, len(arrayOfStrings))
// 2. iterate and convert each element
// for i, v := range arrayOfStrings {
//   arrayOfInts[i] = ....
// }

func main() {
	url := "https://sample-accounts-api.herokuapp.com/users/1"
	// no underscore in golang
	resp, err := http.Get(url)
	// best practice to check and write defer code nearby to original code
	defer resp.Body.Close() // defer here vs derer after error handling is same in this case
	
	if err != nil {
		fmt.Println("error found %v", err) // shadowing error variable... [ from interface error{}] -> avoid naming conflicts with packages
		return
	}
	// return fmt.Errorf("http get %w", err)
	//error handlling missing
	
	// defer resp.Body.Close()
	// defer will be skipped here, but it's ok because an error occurred, so there is no body to close
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error found %v", err)
		return
	}
	
	// var result map[string]any
	// err = json.Unmarshal(body, &result)
	
	// if err != nil {
	// 	fmt.Println("error found")
	// }
	// const attr = "attributes"
	// optimization + safety [ type safety ]
	// attributes, _ := result["attributes"].(map[string]any)
	// attributesJSON, _ := json.Marshal(attributes)
	var user UserResponse  // dont use same var name as struct type name
	err = json.Unmarshal(body, &user)
	
	if err != nil {
		fmt.Println("error found")
	}

	fmt.Println("User: ", user.Attributes.AccountIDs)
	// fmt.Printf("ID: %d, Account IDs: %v\n", user.ID, user.AccountIDs)
	
	// user.Attributes.AccountIDs is slice of ints
	// convert to slice of strings by using strconv.Itoa
	// use strings.Join() to join the elements with ","

}