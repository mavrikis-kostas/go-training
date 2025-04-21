package models

type User struct {
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	AccountIds []int  `json:"account_ids"`
}
