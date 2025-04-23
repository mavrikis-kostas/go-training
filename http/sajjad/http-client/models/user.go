package models

type User struct {
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	AccountIDs []int  `json:"account_ids"`
}
